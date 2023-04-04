from wethenew.wethenewTypes import Thread
from wethenew.backendlink import BackendLinkFlow


def Start():
    t = Thread(BackendLinkFlow, None)
    t.start()
