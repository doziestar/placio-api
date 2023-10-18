from unittest import TestCase
from fastapi.testclient import TestClient
from askai.app.application import create_application


class TestBaseEventHandler(TestCase):
    def test_startup_handler(self):
        app = create_application()
        with self.assertLogs('askai', level='INFO') as cm:

            with TestClient(app):
                pass
            self.assertEqual(cm.output,
                             ['INFO:askai:Starting up ...',
                              'INFO:askai:Shutting down ...'])
