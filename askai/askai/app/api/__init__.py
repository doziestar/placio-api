"""The main APIRouter is defined to include all the sub routers from each
module inside the API folder"""
from fastapi import APIRouter

from .ai import ai_router
from .base import base_router
# TODO: import your modules here.

api_router = APIRouter()
api_router.include_router(base_router, tags=["base"])
api_router.include_router(ai_router, tags=["ai"])

# TODO: include the routers from other modules
