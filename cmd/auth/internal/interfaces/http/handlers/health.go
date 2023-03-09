package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"

	apperrors "github.com/doziestar/tutis-api/pkg/errors"
	httpjson "github.com/doziestar/tutis-api/pkg/http/response/json"

	"google.golang.org/grpc"

	grpcutils "github.com/doziestar/tutis-api/pkg/grpc"
)

// BuildLivenessHandler provides liveness handler
func BuildLivenessHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}

	return http.HandlerFunc(fn)
}

// BuildReadinessHandler provides readiness handler
func BuildReadinessHandler(sqlConn *sql.DB, mongoConn *mongo.Client, connMap map[string]*grpc.ClientConn) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		if sqlConn != nil {
			if err := sqlConn.PingContext(r.Context()); err != nil {
				return apperrors.Wrap(err)
			}
		}
		if mongoConn != nil {
			if err := mongoConn.Ping(r.Context(), nil); err != nil {
				return apperrors.Wrap(err)
			}
		}

		for name, conn := range connMap {
			if !grpcutils.IsConnectionServing(r.Context(), name, conn) {
				return apperrors.New(fmt.Sprintf("gRPC connection %s is not serving", name))
			}
		}

		w.WriteHeader(http.StatusNoContent)

		return nil
	}

	return httpjson.HandlerFunc(fn)
}
