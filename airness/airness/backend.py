from .types import Thread
import time
import requests
import re
from .util import GREEN, BLUE, RESET
from bs4 import BeautifulSoup


def BackendLinkFlow(_, parentThread: Thread):
    print(BLUE + "[*] BackendLinkFlow started..." + RESET)

    headers = {
        "authority": "airness.eu",
        "accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8",
        "accept-language": "en-GB,en-US;q=0.9,en;q=0.8",
        "cache-control": "max-age=0",
        "sec-ch-ua": '"Not_A Brand";v="99", "Brave";v="109", "Chromium";v="109"',
        "sec-ch-ua-mobile": "?0",
        "sec-ch-ua-platform": '"macOS"',
        "sec-fetch-dest": "document",
        "sec-fetch-mode": "navigate",
        "sec-fetch-site": "none",
        "sec-fetch-user": "?1",
        "sec-gpc": "1",
        "upgrade-insecure-requests": "1",
        "user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36",
    }
    session = requests.Session()
    title_shoes = []

    while not parentThread.stop:
        try:
            time.sleep(2)
            try:
                response = session.get(
                    "https://airness.eu/it/uomo/scarpe", headers=headers
                )
                if response.status_code == 200:
                    print(
                        GREEN
                        + "[+] BackendLinkFlow got response {}".format(
                            response.status_code
                        )
                        + RESET
                    )
                    soup = BeautifulSoup(response.text, "html.parser")
                    print(response.text)
                    new_title_shoes = soup.find_all(
                        "h2", {"class": "product-card__title"}
                    )
                    # title_shoes.extend(new_title_shoes)
                    title_shoes.append(new_title_shoes)

                    print(f"\n[*] BackendLinkFlow title_shoes: {new_title_shoes}")
                else:
                    print("\n[!] BackendLinkFlow bad response...")
                    time.sleep(2)
                continue

            except (
                requests.exceptions.ConnectionError
                or requests.exceptions.ConnectTimeout
            ):
                print("\n[!] BackendLinkFlow ConnectionError...")
                continue

        except Exception as e:
            print(f"\n[!] BackendLinkFlow error: {e}")
            continue

    time.sleep(5)
