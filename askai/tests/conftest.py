import pytest
from fastapi.testclient import TestClient
from askai.app.application import create_application
from askai.app.db.session import engine, SessionLocal
from askai.app.db.base import Base
@pytest.fixture
def test_client():
    app = create_application()
    test_client = TestClient(app)
    return test_client


@pytest.fixture
def db_session():
    Base.metadata.create_all(engine)
    session = SessionLocal()
    yield session
    session.close()
    Base.metadata.drop_all(engine)
    engine.dispose()
