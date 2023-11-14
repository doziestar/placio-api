from unittest import TestCase
from fastapi.testclient import TestClient
from africa_insight.app.application import create_application


class TestBaseEventHandler(TestCase):
    def test_startup_handler(self):
        app = create_application()
        with self.assertLogs('africa_insight', level='INFO') as cm:

            with TestClient(app):
                pass
            self.assertEqual(cm.output,
                             ['INFO:africa_insight:Starting up ...',
                              'INFO:africa_insight:Shutting down ...'])
