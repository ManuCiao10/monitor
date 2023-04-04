from oqium.typesOqium import Thread
from oqium.backendlink import BackendLinkFlow


def Start():
    t = Thread(BackendLinkFlow, None)
    t.start()
