import requests
import time
from datetime import datetime


def Start():
    print("[BACKEND] Starting thread...")

    firstRun = True
    ids: list[str] = []

    url = "https://api.nike.com/product_feed/rollup_threads/v2?filter=marketplace%28IT%29&filter=language%28it%29&filter=employeePrice%28true%29&filter=attributeIds%2816633190-45e5-4830-a068-232ac7aea82c%2C193af413-39b0-4d7e-ae34-558821381d3f%2C53e430ba-a5de-4881-8015-68eb1cff459f%29&anchor=0&consumerChannelId=d9a5bc42-4b9c-4976-858a-f159cf99c647&count=60&sort=effectiveStartViewDateDesc"

    while True:
        try:
            try:
                print("[BACKEND] Getting data...")
                response = requests.get(
                    url,
                    headers=headers,
                )
                data = response.json()
                print("[BACKEND] Successful got data.")

            except requests.exceptions.ConnectionError or requests.exceptions.ConnectTimeout:
                print("[BACKEND] Connection error.")
                continue

            except Exception as e:
                print("[BACKEND] Error getting data: {}".format(e))
                continue

            if firstRun:
                id = data["objects"][0]["id"]
                ids.append(id)
                print("[BACKEND] Adding ID (first run): {}".format(id))
                firstRun = False
            else:
                print("[BACKEND] Checking for new IDs...")

                id = data["objects"][0]["id"]
                if id not in ids:
                    ids.append(id)
                    print("[BACKEND] Adding ID: {}".format(id))
                    # get product info

                    status = data["objects"][0]["productInfo"][0]["merchProduct"][
                        "status"
                    ]
                    styleColor = data["objects"][0]["productInfo"][0]["merchProduct"][
                        "styleColor"
                    ]
                    channels = data["objects"][0]["productInfo"][0]["merchProduct"][
                        "channels"
                    ]
                    exclusiveAccess = data["objects"][0]["productInfo"][0][
                        "merchProduct"
                    ]["exclusiveAccess"]
                    try:
                        publishType = data["objects"][0]["productInfo"][0][
                            "merchProduct"
                        ]["publishType"]
                    except:
                        publishType = None

                    imageUrls = data["objects"][0]["productInfo"][0]["imageUrls"][
                        "productImageUrl"
                    ]
                    fullPrice = data["objects"][0]["productInfo"][0]["merchPrice"][
                        "fullPrice"
                    ]
                    available = data["objects"][0]["productInfo"][0]["availability"][
                        "available"
                    ]
                    title = data["objects"][0]["productInfo"][0]["productContent"][
                        "title"
                    ]
                    slug = data["objects"][0]["productInfo"][0]["productContent"][
                        "slug"
                    ]

                    # get launch info
                    try:
                        method = data["objects"][0]["productInfo"][0]["launchView"][
                            "method"
                        ]
                        startEntryDate = data["objects"][0]["productInfo"][0][
                            "launchView"
                        ]["startEntryDate"]
                    except:
                        method = None
                        startEntryDate = None

                    webhook(
                        status=status,
                        styleColor=styleColor,
                        channels=channels,
                        exclusiveAccess=exclusiveAccess,
                        imageUrls=imageUrls,
                        fullPrice=fullPrice,
                        available=available,
                        title=title,
                        slug=slug,
                        method=method,
                        startEntryDate=startEntryDate,
                        publishType=publishType,
                    )

                print("[BACKEND] Sleeping for 300 seconds...")
                time.sleep(300)

        except Exception as e:
            print("[BACKEND] Error: {}".format(e))
            continue


logo = "https://media.discordapp.net/attachments/819084339992068110/1083492784146743416/Screenshot_2023-03-09_at_21.53.23.png"

headers = {
    "authority": "api.nike.com",
    "accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8",
    "accept-language": "en-GB,en-US;q=0.9,en;q=0.8",
    "cache-control": "no-cache",
    "pragma": "no-cache",
    "sec-ch-ua": '"Chromium";v="112", "Brave";v="112", "Not:A-Brand";v="99"',
    "sec-ch-ua-mobile": "?0",
    "sec-ch-ua-platform": '"macOS"',
    "sec-fetch-dest": "document",
    "sec-fetch-mode": "navigate",
    "sec-fetch-site": "none",
    "sec-fetch-user": "?1",
    "sec-gpc": "1",
    "upgrade-insecure-requests": "1",
    "user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36",
}


def webhook(
    status: str,
    styleColor: str,
    channels: list[str],
    exclusiveAccess: bool,
    imageUrls: str,
    fullPrice: int,
    available: bool,
    title: str,
    slug: str,
    method: str,
    startEntryDate: str,
    publishType: str,
):
    print("[BACKEND] Sending webhook...")

    stockx = f"https://stockx.com/search?s={styleColor}"
    klekt = f"https://www.klekt.com/brands?search={styleColor}"
    goat = f"https://www.goat.com/sneakers?query={styleColor}"
    restock = f"https://restocks.net/en/shop/?q={styleColor}"
    cart = "https://www.nike.com/IT/cart"

    webhook = ""
    url = f"https://www.nike.com/it/t/{slug}/{styleColor}"
    embee = []

    array_channels = ""
    for channel in channels:
        array_channels += f"{channel}\n"

    embee.append({"name": "Status", "value": status, "inline": True})
    embee.append({"name": "SKU", "value": styleColor, "inline": True})
    embee.append({"name": "Region", "value": ":flag_it:", "inline": False})
    embee.append(
        {
            "name": "Available at",
            "value": "```\n" + array_channels + "```",
            "inline": False,
        }
    )
    embee.append({"name": "Available", "value": str(available), "inline": True})
    embee.append(
        {"name": "Exclusive Access", "value": str(exclusiveAccess), "inline": True}
    )
    embee.append({"name": "Price", "value": str(fullPrice) + "€", "inline": False})

    if publishType != None:
        embee.append({"name": "Publish type", "value": publishType, "inline": True})

    if method != None:
        embee.append({"name": "Method", "value": method, "inline": True})

    if startEntryDate != None:

        timestamp = datetime.fromisoformat(startEntryDate.replace("Z", "+00:00"))

        formatted_date = timestamp.strftime("%d %B %Y")
        embee.append(
            {"name": "Start Entry Date", "value": formatted_date, "inline": True}
        )
    embee.append(
        {
            "name": "Resell Links",
            "value": f"[StockX]({stockx}) - [Klekt]({klekt}) - [Goat]({goat}) - [Restocks]({restock})",
            "inline": False,
        }
    )
    embee.append({"name": "Useful Links", "value": f"[Cart]({cart})", "inline": False})
    current_time = datetime.now().strftime("%I:%M %p")
    data = {
        "username": "Nike IT",
        "avatar_url": logo,
        "embeds": [
            {
                "title": title,
                "url": url,
                "color": 12298642,
                "thumbnail": {"url": imageUrls},
                "fields": embee,
                "footer": {
                    "text": "Nike by Uzumaki" + " • Today at " + current_time,
                    "icon_url": logo,
                },
            }
        ],
    }

    response = requests.post(webhook, json=data)
    try:
        response.raise_for_status()
        print("[BACKEND] Webhook sent!")
    except requests.exceptions.HTTPError as err:
        print(err)
    except:
        print("[BACKEND] Webhook failed!")
