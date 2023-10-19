"""Main function for running the API service."""
# mypy: ignore-errors
import uvicorn
from app import create_application

app = create_application()

if __name__ == "__main__":
    uvicorn.run("main:app", port=8282, reload=False)  # nosec
