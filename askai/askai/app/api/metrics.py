from typing import Any

from fastapi import APIRouter


metrics_router = APIRouter(prefix="/metrics", tags=["metrics"])


@metrics_router.get("/")
def get_metrics() -> Any:
    return {"message": "Hello, World!"}
