import time
from wethenew.wethenewTypes import Thread
import tls_client
import requests

LOGO = ""

headers = {
    "authority": "api-sell.wethenew.com",
    "accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8",
    "accept-language": "en-GB,en-US;q=0.9,en;q=0.8",
    "cache-control": "no-cache",
    "pragma": "no-cache",
    "sec-ch-ua": '"Chromium";v="110", "Not A(Brand";v="24", "Brave";v="110"',
    "sec-ch-ua-mobile": "?0",
    "sec-ch-ua-platform": '"macOS"',
    "sec-fetch-dest": "document",
    "sec-fetch-mode": "navigate",
    "sec-fetch-site": "none",
    "sec-fetch-user": "?1",
    "sec-gpc": "1",
    "upgrade-insecure-requests": "1",
    "user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36",
}

params = {
    "skip": "0",
    "take": "50",
}


LOGO_UZUMAKI = ""


def webhook_uzumaki(name, image, item):
    url_uzumaki = ""

    field = []

    for sellNow in item["sellNows"]:
        id_sellnow = sellNow["id"]

        value = f"€{sellNow['price']} | {sellNow['size']}"
        size_url = f"http://127.0.0.1:8080/?id={id_sellnow}"

        field.append(
            {
                "name": value,
                "value": "[QUICKTASK]" + "(" + size_url + ")",
                "inline": True,
            }
        )

        data = {
            "username": "WeTheNew Monitor",
            "avatar_url": LOGO_UZUMAKI,
            "embeds": [
                {
                    "title": name,
                    "url": "",
                    "color": 12298642,
                    "description": "> New item to sell",
                    "thumbnail": {"url": image},
                    "fields": field,
                    "footer": {
                        "text": "Powered by Uzumaki",
                        "icon_url": LOGO_UZUMAKI,
                    },
                }
            ],
        }

    result = requests.post(url_uzumaki, json=data)
    try:
        result.raise_for_status()
    except requests.exceptions.HTTPError as err:
        print(err)


def webhook_holding_lab(name, image, item):
    log("sending webhook...")
    url = ""

    field = []

    for sellNow in item["sellNows"]:
        id_ = sellNow["id"]

        value = f"€{sellNow['price']} | {sellNow['size']}"
        size_url = f"https://sell.wethenew.com/instant-sales/{id_}"

        field.append(
            {
                "name": "Payout",
                "value": "[" + value + "]" + "(" + size_url + ")",
                "inline": True,
            }
        )

    data = {
        "username": "WeTheNew",
        "avatar_url": LOGO,
        "embeds": [
            {
                "title": name,
                "url": "",
                "color": 1999236,
                "thumbnail": {"url": image},
                "fields": field,
                "footer": {
                    "text": "WetheNew",
                    "icon_url": LOGO,
                },
            }
        ],
    }

    result = requests.post(url, json=data)
    try:
        result.raise_for_status()
    except requests.exceptions.HTTPError as err:
        print(err)


def log(*args):
    current_time = time.strftime("%H:%M:%S", time.localtime())
    current_time = f"[{current_time}]"
    print(current_time + " [BACKENDLINK]", *args)


def BackendLinkFlow(_, parentThread: Thread):
    log("Starting thread...")

    session = tls_client.Session(client_identifier="chrome_105")
    proxy = ""

    host = proxy.split(":")[0]
    port = proxy.split(":")[1]
    username = proxy.split(":")[2]
    password = proxy.split(":")[3]
    session.proxies = {
        "http": f"http://{username}:{password}@{host}:{port}",
        "https": f"http://{username}:{password}@{host}:{port}",
    }

    firstRun = True
    array = []

    while True:
        try:
            time.sleep(1)  # Delay between each request
            response = session.get(
                "https://api-sell.wethenew.com/sell-nows",
                params=params,
                headers=headers,
            )

            if response.status_code == 200:
                log("Success Getting Data")
                data = response.json()

                try:
                    if firstRun:
                        log("adding sell nows...")
                        firstRun = False

                        result = data["results"]
                        for item in result:
                            array.append(item["id"])

                    else:
                        log("no new sell nows...")
                        result = data["results"]
                        # print(array) For testing purpose
                        # array.remove(97) For testing purpose
                        for item in result:
                            if item["id"] not in array:
                                log("new sell now")

                                array.append(item["id"])
                                webhook_holding_lab(item["name"], item["image"], item)
                                webhook_uzumaki(item["name"], item["image"], item)

                                time.sleep(5)

                except Exception as e:
                    print("Error Getting Data" + str(e))
                    time.sleep(5)

            else:
                log("Error Getting Data")
                time.sleep(10)

        except Exception as e:
            print("Error " + str(e))
            time.sleep(10)
