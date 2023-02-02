from tba.types import Thread
from tba.backendlink import BackendLinkFlow

# from dsm.keyword import KeywordFlow


# women https://www.the-broken-arm.com/en/58-sneakers
# men https://www.the-broken-arm.com/en/57-sneakers
# xml https://www.the-broken-arm.com/1_en_0_sitemap.xml


def Start():
    t = Thread(BackendLinkFlow, None)
    t.start()

    # keywordThread = Thread(KeywordFlow, t)
    # keywordThread.start()
