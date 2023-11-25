package smartMenu

import (
	"context"
	"fmt"
	"image/color"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"placio-app/domains/media"
	"placio-app/ent"
	"placio-app/ent/category"
	"placio-app/ent/menu"
	"placio-app/ent/menuitem"
	"placio-app/ent/place"
	"placio-app/ent/placetable"
	"placio-pkg/errors"
	"strings"
	"time"

	firebase "firebase.google.com/go"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"google.golang.org/api/option"

	"github.com/cloudinary/cloudinary-go/v2"
	qrcode "github.com/skip2/go-qrcode"

	"github.com/google/uuid"
)

type SmartMenuService struct {
	client       *ent.Client
	mediaService media.MediaService
	cloud        *cloudinary.Cloudinary
}

type ISmartMenu interface {
	CreateMenu(context.Context, string, string, *ent.Menu, []*multipart.FileHeader) (*ent.Menu, error)
	GetMenus(context.Context, string) ([]*ent.Menu, error)
	GetMenuByID(context.Context, string) (*ent.Menu, error)
	UpdateMenu(context.Context, string, string, *ent.Menu) (*ent.Menu, error)
	DeleteMenu(context.Context, string) error
	RestoreMenu(context.Context, string) (*ent.Menu, error)

	CreateMenuItem(context.Context, string, *ent.MenuItem, []*multipart.FileHeader) (*ent.MenuItem, error)
	GetMenuItems(context.Context, string) ([]*ent.MenuItem, error)
	GetMenuItemByID(context.Context, string) (*ent.MenuItem, error)
	UpdateMenuItem(context.Context, string, *ent.MenuItem) (*ent.MenuItem, error)
	DeleteMenuItem(context.Context, string) error
	RestoreMenuItem(context.Context, string) (*ent.MenuItem, error)

	CreateTable(context.Context, string, string, *ent.PlaceTable) (*ent.PlaceTable, error)
	GetTables(context.Context, string) ([]*ent.PlaceTable, error)
	GetTableByID(context.Context, string) (*ent.PlaceTable, error)
	UpdateTable(context.Context, string, string, *ent.PlaceTable) (*ent.PlaceTable, error)
	DeleteTable(context.Context, string) error
	RestoreTable(context.Context, string) (*ent.PlaceTable, error)
	RegenerateQRCode(context.Context, string) (*ent.PlaceTable, error)
}

func NewSmartMenuService(client *ent.Client, mediaService media.MediaService, cloud *cloudinary.Cloudinary) ISmartMenu {
	return &SmartMenuService{client: client, mediaService: mediaService, cloud: cloud}
}

func (s *SmartMenuService) CreateMenuItem(ctx context.Context, menuId string, menuItemDto *ent.MenuItem, mediaFiles []*multipart.FileHeader) (*ent.MenuItem, error) {
	existingMenuItem, err := s.client.MenuItem.
		Query().
		Where(menuitem.Name(strings.ToLower(menuItemDto.Name))).
		Where(menuitem.HasMenuWith(menu.ID(menuId))).
		Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		log.Println("Error checking if menu item with name", menuItemDto.Name, "exists for menu", menuId, ":", err)
		return nil, err
	}
	if existingMenuItem != nil {
		log.Println("Menu item with name", menuItemDto.Name, "already exists for menu", menuId)
		return existingMenuItem, nil
	}

	log.Println("Creating menu item with name", menuItemDto.Name)

	// Create and return the new menu item
	newMenuItem, err := s.client.MenuItem.
		Create().
		SetID(uuid.New().String()).
		SetName(menuItemDto.Name).
		SetDescription(menuItemDto.Description).
		SetPrice(menuItemDto.Price).
		SetIsAvailable(menuItemDto.IsAvailable).
		SetAvailableFrom(menuItemDto.AvailableFrom).
		SetOptions(menuItemDto.Options).
		SetPrice(menuItemDto.Price).
		SetIsNew(true).
		AddMenuIDs(menuId).
		Save(ctx)
	if err != nil {
		log.Println("Error creating menu item with name", menuItemDto.Name, ":", err)
		return nil, err
	}

	// Create the media for the menu item
	if len(mediaFiles) > 0 {
		go func(menuItemID string, mediaFiles []*multipart.FileHeader) {
			ctx := context.Background() // Use a background context for the asynchronous operation
			log.Println("Uploading media for menu item with ID", menuItemID)

			media, err := s.mediaService.UploadAndCreateMedia(ctx, mediaFiles)
			if err != nil {
				log.Println("Error uploading media for menu item with ID", menuItemID, ":", err)
				return
			}

			_, err = s.client.MenuItem.
				UpdateOneID(menuItemID).
				AddMedia(media...).
				Save(ctx)
			if err != nil {
				log.Println("Error adding media to menu item with ID", menuItemID, ":", err)
				return
			}

			log.Println("Media uploaded and added to menu item with ID", menuItemID)
		}(newMenuItem.ID, mediaFiles)
	}

	return newMenuItem, nil
}

func (s *SmartMenuService) GetMenuItems(ctx context.Context, menuId string) ([]*ent.MenuItem, error) {
	log.Println("Getting menu items for menu with ID", menuId)
	return s.client.MenuItem.
		Query().
		Where(menuitem.HasMenuWith(menu.ID(menuId))).
		WithMedia().
		All(ctx)
}

func (s *SmartMenuService) GetMenuItemByID(ctx context.Context, menuItemId string) (*ent.MenuItem, error) {
	return s.client.MenuItem.
		Query().
		Where(menuitem.ID(menuItemId)).
		WithMedia().
		Only(ctx)
}

func (s *SmartMenuService) UpdateMenuItem(ctx context.Context, menuItemId string, menuItemDto *ent.MenuItem) (*ent.MenuItem, error) {
	update := s.client.MenuItem.UpdateOneID(menuItemId)

	if menuItemDto.Name != "" {
		update.SetName(menuItemDto.Name)
	}
	if menuItemDto.Price != 0 {
		update.SetPrice(menuItemDto.Price)
	}
	if menuItemDto.IsAvailable != true {
		update.SetIsAvailable(menuItemDto.IsAvailable)
	}

	// Optional string fields
	if menuItemDto.Description != "" {
		update.SetDescription(menuItemDto.Description)
	}
	if menuItemDto.Currency != "" {
		update.SetCurrency(menuItemDto.Currency)
	}
	if menuItemDto.DeletedAt != "" {
		update.SetDeletedAt(menuItemDto.DeletedAt)
	}
	if menuItemDto.ImageURL != "" {
		update.SetImageURL(menuItemDto.ImageURL)
	}
	if menuItemDto.ChefSpecialNote != "" {
		update.SetChefSpecialNote(menuItemDto.ChefSpecialNote)
	}
	if menuItemDto.Category != "" {
		update.SetCategory(menuItemDto.Category)
	}

	if menuItemDto.Season != "" {
		update.SetSeason(menuItemDto.Season)
	}
	if menuItemDto.PromotionDescription != "" {
		update.SetPromotionDescription(menuItemDto.PromotionDescription)
	}

	// Optional int fields
	if menuItemDto.PreparationTime != 0 {
		update.SetPreparationTime(menuItemDto.PreparationTime)
	}
	if menuItemDto.Calories != 0 {
		update.SetCalories(menuItemDto.Calories)
	}
	if menuItemDto.ServeSize != 0 {
		update.SetServeSize(menuItemDto.ServeSize)
	}

	if menuItemDto.DiscountPercentage != 0 {
		update.SetDiscountPercentage(menuItemDto.DiscountPercentage)
	}

	// Optional boolean fields
	update.SetIsDeleted(menuItemDto.IsDeleted)
	update.SetIsFeatured(menuItemDto.IsFeatured)
	update.SetIsNew(menuItemDto.IsNew)
	update.SetIsSeasonal(menuItemDto.IsSeasonal)

	if menuItemDto.Options != nil {
		update.SetOptions(menuItemDto.Options)
	}
	if menuItemDto.Allergens != nil {
		update.SetAllergens(menuItemDto.Allergens)
	}
	if menuItemDto.Tags != nil {
		update.SetTags(menuItemDto.Tags)
	}

	return update.Save(ctx)
}

func (s *SmartMenuService) DeleteMenuItem(ctx context.Context, menuItemId string) error {
	return s.client.MenuItem.
		DeleteOneID(menuItemId).
		Exec(ctx)
}

func (s *SmartMenuService) RestoreMenuItem(ctx context.Context, menuItemId string) (*ent.MenuItem, error) {
	// TODO: Implement restore functionality if soft delete is implemented
	return nil, nil
}

func (s *SmartMenuService) CreateMenu(ctx context.Context, placeId, userId string, menuDto *ent.Menu, medias []*multipart.FileHeader) (*ent.Menu, error) {

	// Check if the menu already exists in the database for this place
	existingMenu, err := s.client.Menu.
		Query().
		Where(menu.Name(strings.ToLower(menuDto.Name))).
		Where(menu.HasPlaceWith(place.ID(placeId))).
		Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		log.Println("Error checking if menu with name", menuDto.Name, "exists for place", placeId, ":", err)
		return nil, err
	}
	if existingMenu != nil {
		log.Println("Menu with name", menuDto.Name, "already exists for place", placeId)
		return existingMenu, nil
	}

	log.Println("Creating menu with name", menuDto.Name)

	// Start creating the menu
	menuCreate := s.client.Menu.
		Create().
		SetID(uuid.New().String()).
		SetName(menuDto.Name).
		SetDescription(menuDto.Description).
		SetOptions(menuDto.Options).
		SetMenuItemType(menuDto.MenuItemType).
		SetIsAvailable(menuDto.IsAvailable).
		AddPlaceIDs(placeId).
		AddCreatedByIDs(userId)

	// Set the food type and dietary type if menuItemType is food
	if menuDto.MenuItemType == menu.MenuItemType("food") {
		if menuDto.FoodType == "" {
			return nil, errors.New("foodType must be provided when menuItemType is 'food'")
		}
		if menuDto.DietaryType == "" {
			return nil, errors.New("dietaryType must be provided when menuItemType is 'food'")
		}
		menuCreate = menuCreate.SetFoodType(menuDto.FoodType).SetDietaryType(menuDto.DietaryType)
	}

	// Set the drink type if menuItemType is drink
	if menuDto.MenuItemType == menu.MenuItemType("drink") && menuDto.DrinkType != "" {
		menuCreate = menuCreate.SetDrinkType(menuDto.DrinkType)
	}

	// Save the new menu
	newMenu, err := menuCreate.Save(ctx)
	if err != nil {
		log.Println("Error creating menu with name", menuDto.Name, ":", err)
		return nil, err
	}

	// Handle media files asynchronously if present
	if len(medias) > 0 {
		newCtx := context.Background() // Use a background context for the asynchronous operation
		go s.handleMenuMedia(newCtx, newMenu.ID, medias)
	}

	return newMenu, nil
}

func (s *SmartMenuService) handleMenuMedia(ctx context.Context, menuID string, medias []*multipart.FileHeader) {
	log.Println("Uploading media for menu with ID", menuID)

	media, err := s.mediaService.UploadAndCreateMedia(ctx, medias)
	if err != nil {
		log.Println("Error uploading media for menu with ID", menuID, ":", err)
		return
	}

	_, err = s.client.Menu.
		UpdateOneID(menuID).
		AddMedia(media...).
		Save(ctx)
	if err != nil {
		log.Println("Error adding media to menu with ID", menuID, ":", err)
		return
	}

	log.Println("Media uploaded and added to menu with ID", menuID)
}

func (s *SmartMenuService) ensureCategoryExists(ctx context.Context, categoryName string) (*ent.Category, error) {
	catData, err := s.client.Category.
		Query().
		Where(category.Name(categoryName)).
		Only(ctx)
	if ent.IsNotFound(err) {
		return s.client.Category.
			Create().
			SetName(categoryName).
			Save(ctx)
	}
	return catData, err
}

func (s *SmartMenuService) GetMenus(ctx context.Context, placeId string) ([]*ent.Menu, error) {
	return s.client.Menu.
		Query().
		Where(menu.HasPlaceWith(place.ID(placeId))).
		Where(menu.IsDeleted(false)).
		WithMedia().
		WithPlace(func(q *ent.PlaceQuery) {
			q.WithBusiness(func(bq *ent.BusinessQuery) {
			})
			q.WithMedias()
		}).All(ctx)
}

func (s *SmartMenuService) GetMenuByID(ctx context.Context, menuId string) (*ent.Menu, error) {
	return s.client.Menu.
		Query().
		Where(menu.ID(menuId)).
		WithMedia().
		WithPlace(func(q *ent.PlaceQuery) {
			q.WithBusiness()
		}).
		WithCategories(func(q *ent.CategoryQuery) {
		}).
		Only(ctx)
}

func (s *SmartMenuService) UpdateMenu(ctx context.Context, menuId, userId string, menuData *ent.Menu) (*ent.Menu, error) {
	// Validate input (example: check for non-empty name)
	if menuData.Name == "" {
		return nil, errors.New("menu name cannot be empty")
	}

	// Prepare update operation
	updateOp := s.client.Menu.UpdateOneID(menuId)

	if menuData.Name != "" {
		updateOp = updateOp.SetName(menuData.Name)
	}
	if menuData.Description != "" {
		updateOp = updateOp.SetDescription(menuData.Description)
	}

	updateOp = updateOp.SetUpdatedAt(time.Now().Local()).AddUpdatedByIDs(userId)

	// Update common fields
	updateOp = updateOp.SetName(menuData.Name).
		SetDescription(menuData.Description).
		SetOptions(menuData.Options).
		SetIsAvailable(menuData.IsAvailable).
		SetUpdatedAt(time.Now().Local()).
		AddUpdatedByIDs(userId)

	// Handle menu type-specific updates
	if menuData.MenuItemType != "" {
		switch menuData.MenuItemType {
		case menu.MenuItemType("food"):
			if menuData.FoodType == "" || menuData.DietaryType == "" {
				return nil, errors.New("both foodType and dietaryType must be provided for 'food' menuItemType")
			}
			updateOp = updateOp.SetFoodType(menuData.FoodType).SetDietaryType(menuData.DietaryType)

		case menu.MenuItemType("drink"):
			if menuData.DrinkType == "" {
				return nil, errors.New("drinkType must be provided for 'drink' menuItemType")
			}
			updateOp = updateOp.SetDrinkType(menuData.DrinkType)
		}
	}

	// Perform the update
	updatedMenu, err := updateOp.Save(ctx)
	if err != nil {
		log.Printf("Error updating menu with ID %s: %v", menuId, err)
		return nil, err
	}

	return updatedMenu, nil
}

func (s *SmartMenuService) DeleteMenu(ctx context.Context, menuId string) error {
	return s.client.Menu.
		UpdateOneID(menuId).
		SetDeletedAt(time.Now().String()).
		SetIsDeleted(true).
		Exec(ctx)
}

func (s *SmartMenuService) RestoreMenu(ctx context.Context, menuId string) (*ent.Menu, error) {
	return s.client.Menu.
		UpdateOneID(menuId).
		ClearDeletedAt().
		SetIsDeleted(false).
		Save(ctx)
}

func (s *SmartMenuService) CreateTable(ctx context.Context, placeId, userId string, table *ent.PlaceTable) (*ent.PlaceTable, error) {
	log.Println("Creating table with number", table)
	tableInfo, err := s.client.PlaceTable.
		Create().
		SetID(uuid.New().String()).
		SetNumber(table.Number).
		SetCapacity(table.Capacity).
		SetName(table.Name).
		SetType(table.Type).
		SetIsPremium(table.IsPremium).
		SetIsVip(table.IsVip).
		SetIsActive(table.IsActive).
		SetIsReserved(table.IsReserved).
		SetCreatedByID(userId).
		SetPlaceID(placeId).
		Save(ctx)

	if err != nil {
		log.Println("Error creating table with number", table.Number, ":", err)
		return nil, err
	}

	return tableInfo, nil
}

func (s *SmartMenuService) GetTables(ctx context.Context, placeId string) ([]*ent.PlaceTable, error) {
	return s.client.Place.
		Query().
		Where(place.ID(placeId)).
		Where(place.HasTables()).
		QueryTables().
		Where(placetable.IsDeleted(false)).
		All(ctx)
}

func (s *SmartMenuService) GetTableByID(ctx context.Context, tableId string) (*ent.PlaceTable, error) {
	return s.client.PlaceTable.
		Query().
		Where(placetable.ID(tableId)).
		Only(ctx)
}

func (s *SmartMenuService) UpdateTable(ctx context.Context, tableId, userId string, table *ent.PlaceTable) (*ent.PlaceTable, error) {
	return s.client.PlaceTable.
		UpdateOneID(tableId).
		SetNumber(table.Number).
		SetNumber(table.Number).
		SetCapacity(table.Capacity).
		SetName(table.Name).
		SetType(table.Type).
		SetIsPremium(table.IsPremium).
		SetIsVip(table.IsVip).
		SetIsActive(table.IsActive).
		SetIsReserved(table.IsReserved).
		SetUpdatedByID(userId).
		Save(ctx)
}

func (s *SmartMenuService) DeleteTable(ctx context.Context, tableId string) error {
	return s.client.PlaceTable.
		UpdateOneID(tableId).
		SetDeletedAt(time.Now().String()).
		SetIsDeleted(true).
		Exec(ctx)
}

func (s *SmartMenuService) RestoreTable(ctx context.Context, tableId string) (*ent.PlaceTable, error) {
	return s.client.PlaceTable.
		UpdateOneID(tableId).
		ClearDeletedAt().
		SetIsDeleted(false).
		Save(ctx)
}

func (s *SmartMenuService) RegenerateQRCode(ctx context.Context, tableId string) (*ent.PlaceTable, error) {
	table, err := s.client.PlaceTable.
		Query().
		Where(placetable.IDEQ(tableId)).
		WithPlace(func(pq *ent.PlaceQuery) {
			pq.WithBusiness(func(bq *ent.BusinessQuery) {
				bq.WithWebsites()
			})
		}).
		Only(ctx)
	if err != nil {
		log.Println("Error fetching table with ID", tableId, ":", err)
		return nil, fmt.Errorf("failed querying place table: %w", err)
	}

	url := fmt.Sprintf("https://placio.io/%s/menus/?table=%d&placeId=%s",
		table.Edges.Place.Edges.Business.Edges.Websites.DomainName,
		table.Number,
		table.Edges.Place.ID,
	)

	qr, err := qrcode.New(url, qrcode.Medium)
	if err != nil {
		log.Println("Error generating QR code:", err)
		return nil, fmt.Errorf("failed to generate QR code: %w", err)
	}

	qr.ForegroundColor = color.RGBA{R: 139, G: 0, B: 0, A: 255}
	qr.BackgroundColor = color.White

	png, err := qr.PNG(256)
	if err != nil {
		log.Println("Error converting QR code to PNG:", err)
		return nil, fmt.Errorf("failed to convert QR code to PNG: %w", err)
	}

	tmpFile, err := ioutil.TempFile("", "qr-code-*.png")
	if err != nil {
		return nil, fmt.Errorf("failed to create temporary file: %w", err)
	}

	_, err = tmpFile.Write(png)
	if err != nil {
		tmpFile.Close()
		return nil, fmt.Errorf("failed to write to temporary file: %w", err)
	}
	tmpFile.Close()

	//signedURL, err := s.uploadQRCodeToFirebase(ctx, tmpFile.Name(), "image/png")
	signedURL, err := s.uploadQRCodeToDigitalOceanSpace(ctx, table, tmpFile.Name(), "image/png")
	if err != nil {
		return nil, err
	}

	os.Remove(tmpFile.Name())

	updatedTable, err := s.client.PlaceTable.
		UpdateOneID(table.ID).
		SetQrCode(signedURL).
		Save(ctx)
	if err != nil {
		log.Println("Error updating table with ID", tableId, ":", err)
		return nil, fmt.Errorf("failed updating place table: %w", err)
	}

	return updatedTable, nil
}

func (s *SmartMenuService) uploadQRCodeToDigitalOceanSpace(ctx context.Context, table *ent.PlaceTable, filePath, contentType string) (string, error) {
	// DigitalOcean Spaces credentials and configuration
	accessKeyID := "DO00YJ68Y7KMTYP3J7HE"
	secretAccessKey := "P55ReutOGyn1d4qThoPCMj+O7qCUggr/Y+DQIUwYtjc"
	spaceName := "placio"
	endpoint := "https://placio.fra1.digitaloceanspaces.com"
	cdnEndpoint := "https://placio.fra1.cdn.digitaloceanspaces.com"

	// Create a new session without specifying a region
	sess, err := session.NewSession(&aws.Config{
		Endpoint:         aws.String(endpoint),
		Region:           aws.String("fra1"),
		Credentials:      credentials.NewStaticCredentials(accessKeyID, secretAccessKey, ""),
		S3ForcePathStyle: aws.Bool(true),
	})
	if err != nil {
		return "", fmt.Errorf("error creating session: %w", err)
	}

	// Open the file to upload
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Construct the unique file name
	domainName := table.Edges.Place.Edges.Business.Edges.Websites.DomainName
	tableNumber := table.Number
	uniqueFileName := fmt.Sprintf("barcode/%s-%s-%s", domainName, string(tableNumber), filepath.Base(filePath))

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	// Upload the file and get the URL in return
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:             aws.String(spaceName),
		Key:                aws.String(uniqueFileName),
		ACL:                aws.String("public-read"),
		Body:               file,
		ContentType:        aws.String(contentType),
		ContentDisposition: aws.String("inline"),
	})

	if err != nil {
		return "", fmt.Errorf("failed to upload to DigitalOcean Space: %w", err)
	}

	// Direct URL of the uploaded file
	directURL := result.Location

	// Construct the CDN URL by replacing the direct URL's domain with the CDN endpoint
	cdnURL := strings.Replace(directURL, endpoint, cdnEndpoint, 1)

	return cdnURL, nil
}

func (s *SmartMenuService) uploadQRCodeToFirebase(ctx context.Context, filePath, contentType string) (string, error) {
	conf := &firebase.Config{StorageBucket: "placio-383019.appspot.com"}
	opt := option.WithCredentialsFile("serviceAccount.json")
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		return "", fmt.Errorf("error initializing firebase app: %w", err)
	}

	client, err := app.Storage(ctx)
	if err != nil {
		return "", fmt.Errorf("error getting Storage client: %w", err)
	}

	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	bucket, err := client.Bucket("placio-383019.appspot.com")
	if err != nil {
		return "", fmt.Errorf("error getting default bucket: %w", err)
	}

	object := bucket.Object("placio/" + filepath.Base(filePath))
	wc := object.NewWriter(ctx)
	if _, err = io.Copy(wc, file); err != nil {
		return "", fmt.Errorf("error writing to Firebase Storage: %w", err)
	}
	if err = wc.Close(); err != nil {
		return "", fmt.Errorf("error closing writer: %w", err)
	}

	signedURL, err := media.GenerateSignedURL(ctx, "placio-383019.appspot.com", "placio/"+filepath.Base(filePath))
	if err != nil {
		return "", fmt.Errorf("failed to generate signed URL: %w", err)
	}

	return signedURL, nil
}
