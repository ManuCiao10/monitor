import requests
import time

dict_size = {
    "4723523": "34.5",
    "4723524": "35",
    "4723525": "36",
    "4723526": "36.5",
    "4723527": "37",
    "4723528": "38",
    "4723529": "38.5",
    "4723530": "39",
    "4723531": "40",
    "4723532": "40.5",
    "4723533": "41",
    "4723534": "42",
    "4723535": "42.5",
    "4723536": "43",
    "4723537": "44",
    "4723538": "44.5",
    "4723539": "45",
    "4723540": "46",
    "4723541": "47",
}


def webhook(dict_stock: dict):
    price = "90 €"
    name = "SCARPE KNU SKOOL"
    link = "https://www.vans.it/shop/it/vans-it/scarpe-knu-skool-vn0009qc6bt"
    img = "https://s7d2.scene7.com/is/image/VansEU/VN0009QC6BT-HERO"
    Logo = "https://media.discordapp.net/attachments/819084339992068110/1083492784146743416/Screenshot_2023-03-09_at_21.53.23.png"
    webhook_uzumaki = ""
    embee = []
    content = ""

    for key in dict_stock:
        content += f"{key} - [{dict_stock[key]}]\n"

    embee.append(
        {
            "name": "Sizes",
            "value": content,
            "inline": True,
        }
    )

    embee.append({"name": "Price", "value": price, "inline": False})
    data = {
        "username": "Vans Monitor",
        "avatar_url": Logo,
        "embeds": [
            {
                "title": name,
                "url": link,
                "color": 12298642,
                "description": "> Stock updated",
                "thumbnail": {"url": img},
                "fields": embee,
                "footer": {
                    "text": "Vans Monitor | Uzumaki",
                    "icon_url": Logo,
                },
            }
        ],
    }

    result = requests.post(webhook_uzumaki, json=data)
    try:
        result.raise_for_status()
        print("Webhook delivered successfully, code {}.".format(result.status_code))
    except requests.exceptions.HTTPError as err:
        print(err)


def Start():
    print("Starting monitor vans...")

    firstRun = True
    uno = {}

    headers = {
        "authority": "www.vans.it",
        "accept": "*/*",
        "accept-language": "en-GB,en-US;q=0.9,en;q=0.8",
        "cache-control": "no-cache",
        "pragma": "no-cache",
        "referer": "https://www.vans.it/shop/it/vans-it/scarpe-knu-skool-vn0009qc6bt",
        "sec-ch-ua": '"Chromium";v="112", "Brave";v="112", "Not:A-Brand";v="99"',
        "sec-ch-ua-mobile": "?0",
        "sec-ch-ua-platform": '"macOS"',
        "sec-fetch-dest": "empty",
        "sec-fetch-mode": "cors",
        "sec-fetch-site": "same-origin",
        "sec-gpc": "1",
        "user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36",
        "x-requested-with": "XMLHttpRequest",
    }

    params = {
        "requestype": "ajax",
        "storeId": "10161",
        "langId": "-4",
        "productId": "4721629",
        "requesttype": "ajax",
    }

    while True:
        try:
            try:
                response = requests.get(
                    "https://www.vans.it/webapp/wcs/stores/servlet/VFAjaxProductAvailabilityView",
                    params=params,
                    headers=headers,
                )

            except requests.exceptions.ConnectionError or requests.exceptions.ConnectTimeout:
                print("Connection error...")
                continue

            if response.status_code == 200:
                if firstRun:
                    print("Adding (firstrun) product: ")

                    try:
                        data = response.json()
                        firstRun = False
                    except Exception as e:
                        print("Error getting data: ", e)
                        time.sleep(5)
                        continue
                else:
                    print("Checking for new products: ")
                    try:
                        dataNew = response.json()
                        print("Success got data")
                    except Exception as e:
                        print("Error getting data: ", e)
                        time.sleep(5)
                        continue

                    if data != dataNew:
                        print("New product found!")
                        data = dataNew

                        dict_stock = data["stock"]

                        for key in dict_size:
                            euro_size = str(dict_size[key])
                            uno[euro_size] = dict_stock[key]

                        webhook(dict_stock=uno)

                    else:
                        print("No new products found")
                        time.sleep(120)

            else:
                print("Error getting data: ", response.status_code)
                time.sleep(5)
                continue
        except Exception as e:
            print("[KEYWORD] Error: {}".format(e))
            time.sleep(5)
            continue
