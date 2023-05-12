import {
  IsEmail,
  IsString,
  IsNotEmpty,
  MinLength,
  MaxLength,
  IsOptional,
  IsBoolean,
  IsUUID,
  ValidateNested,
  IsInstance,
  IsPhoneNumber,
  IsEnum,
} from 'class-validator';

// AccountType Enum
export enum AccountType {
  USER = 'user',
  BUSINESS = 'business',
}

// CreateUserDto
export class CreateUserDto {
  @IsEmail()
  public email: string;

  @IsString()
  @IsNotEmpty()
  public name: string;

  @IsString()
  @IsNotEmpty()
  public username: string;

  @IsString()
  @IsOptional()
  @IsPhoneNumber()
  public phone?: string;

  @IsEnum(AccountType)
  public account_type: AccountType;

  @IsString()
  @IsNotEmpty()
  @MinLength(9)
  @MaxLength(32)
  public password: string;

  @IsString()
  @IsNotEmpty()
  @MinLength(9)
  @MaxLength(32)
  public confirm_password: string;
}

// UpdateUserDto
export class UpdateUserDto {
  @IsOptional()
  @IsString()
  @IsNotEmpty()
  public name?: string;

  @IsOptional()
  @IsString()
  @IsNotEmpty()
  @MinLength(9)
  @MaxLength(32)
  public password?: string;

  @IsOptional()
  @IsBoolean()
  public support_enabled?: boolean;

  @IsOptional()
  @IsBoolean()
  public '2fa_enabled'?: boolean;
}

// LoginUserDto
export class LoginUserDto {
  @IsEmail()
  @IsOptional()
  public email: string;

  @IsString()
  @IsOptional()
  public username: string;

  @IsString()
  @IsNotEmpty()
  @MinLength(9)
  @MaxLength(32)
  public password: string;
}

// Enable2FADto
export class Enable2FADto {
  @IsBoolean()
  public '2fa_enabled': boolean;

  @IsOptional()
  @IsString()
  @IsNotEmpty()
  public '2fa_secret'?: string;

  @IsOptional()
  @IsString()
  @IsNotEmpty()
  public '2fa_backup_code'?: string;
}

// UpdateSocialMediaDto
export class UpdateSocialMediaDto {
  @IsOptional()
  @IsString()
  @IsNotEmpty()
  public facebook_id?: string;

  @IsOptional()
  @IsString()
  @IsNotEmpty()
  public twitter_id?: string;

  @IsOptional()
  @ValidateNested()
  @IsInstance(Object)
  public twitter?: {
    accessToken?: string;
    refreshToken?: string;
    userId?: string;
    userName?: string;
    codeVerifier?: string;
    state?: string;
    name?: string;
    dateCreated?: Date;
    expiresIn?: Date;
  };

  @IsOptional()
  @ValidateNested()
  @IsInstance(Object)
  public google?: {
    accessToken?: string;
    refreshToken?: string;
    userId?: string;
    email?: string;
    dateCreated?: Date;
  };
}

// UpdateGeneralSettingsDto
export class UpdateGeneralSettingsDto {
  @IsUUID()
  public generalSettings: string;
}
