"""Module containing FastAPI instance related functions and classes."""
# mypy: ignore-errors
import logging.config
from fastapi import FastAPI
from starlette.middleware.base import BaseHTTPMiddleware
from .api import api_router
from .version import __version__


def create_application() -> FastAPI:
    """Create a FastAPI instance.

    Returns:
        object of FastAPI: the fastapi application instance.
    """
    # settings = get_settings()
    application = FastAPI(title="AskAI API",
                          debug=True,
                          version=__version__,
                          openapi_url="/api/v1/openapi.json")

    # Set all CORS enabled origins
    # if settings.CORS_ORIGINS:
    #     application.add_middleware(
    #         CORSMiddleware,
    #         allow_origins=[str(origin) for origin in
    #                        settings.CORS_ORIGINS],
    #         allow_origin_regex=settings.CORS_ORIGIN_REGEX,
    #         allow_credentials=settings.CORS_CREDENTIALS,
    #         allow_methods=settings.CORS_METHODS,
    #         allow_headers=settings.CORS_HEADERS,
    #     )

    # add defined routers
    application.include_router(api_router, prefix="/api/v1")

    # logging.config.dictConfig(settings.LOGGING_CONFIG)

    return application
