from datetime import datetime
import requests


def webhook_titan(
    data: dict,
):
    print("[BACKEND] Sending webhook...")

    stockx = f"https://stockx.com/search?s={data['styleColor']}"
    klekt = f"https://www.klekt.com/brands?search={data['styleColor']}"
    goat = f"https://www.goat.com/sneakers?query={data['styleColor']}"
    restock = f"https://restocks.net/en/shop/?q={data['styleColor']}"
    cart = "https://www.nike.com/IT/cart"
    app = "http://atc.yeet.ai/redirect?link=mynike://x-callback-url/product-details?style-color={data['styleColor']}&redirect=true"
    url = f"https://www.nike.com/it/t/{data['slug']}/{data['styleColor']}"

    webhook = ""
    logo = ""

    embee = []

    array_channels = ""
    for channel in data["channels"]:
        array_channels += f"{channel}\n"

    embee.append({"name": "Status", "value": data["status"], "inline": True})
    embee.append({"name": "SKU", "value": data["styleColor"], "inline": True})
    embee.append({"name": "Region", "value": ":flag_it:", "inline": False})
    embee.append(
        {
            "name": "Available at",
            "value": "```\n" + array_channels + "```",
            "inline": False,
        }
    )
    embee.append({"name": "Available", "value": str(data["available"]), "inline": True})
    embee.append(
        {"name": "Price", "value": str(data["fullPrice"]) + "€", "inline": False}
    )

    # based of the size of the line data["size"]

    line = len(data["size"].splitlines())
    number_of_element = line / 5
    number_of_element = int(number_of_element)

    for i in range(number_of_element):
        # create a list of 5 elements
        list_of_5 = data["size"].splitlines()[i * 5 : (i + 1) * 5]
        # join the list of 5 elements
        list_of_5 = "\n".join(list_of_5)

        embee.append({"name": "Size", "value": list_of_5, "inline": True})

    embee.append(
        {
            "name": "Resell Links",
            "value": f"[StockX]({stockx}) - [Klekt]({klekt}) - [Goat]({goat}) - [Restocks]({restock})",
            "inline": False,
        }
    )
    embee.append(
        {
            "name": "Useful Links",
            "value": f"[Cart]({cart}) - [App]({app})",
            "inline": False,
        }
    )
    current_time = datetime.now().strftime("%I:%M %p")

    data = {
        "username": "Nike Restock IT",
        "avatar_url": logo,
        "embeds": [
            {
                "title": data["title"],
                "url": url,
                "color": 12058624,
                "thumbnail": {"url": data["imageUrls"]},
                "fields": embee,
                "footer": {
                    "text": "Nike by Titan" + " • Today at " + current_time,
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


def webhook_uzumaki(
    data: dict,
):
    print("[BACKEND] Sending webhook...")

    stockx = f"https://stockx.com/search?s={data['styleColor']}"
    klekt = f"https://www.klekt.com/brands?search={data['styleColor']}"
    goat = f"https://www.goat.com/sneakers?query={data['styleColor']}"
    restock = f"https://restocks.net/en/shop/?q={data['styleColor']}"
    cart = "https://www.nike.com/IT/cart"
    app = "http://atc.yeet.ai/redirect?link=mynike://x-callback-url/product-details?style-color={data['styleColor']}&redirect=true"

    url = f"https://www.nike.com/it/t/{data['slug']}/{data['styleColor']}"

    webhook = ""
    logo_uzumaki = "https://media.discordapp.net/attachments/819084339992068110/1083492784146743416/Screenshot_2023-03-09_at_21.53.23.png"

    embee = []

    array_channels = ""
    for channel in data["channels"]:
        array_channels += f"{channel}\n"

    embee.append({"name": "Status", "value": data["status"], "inline": True})
    embee.append({"name": "SKU", "value": data["styleColor"], "inline": True})
    embee.append({"name": "Region", "value": ":flag_it:", "inline": False})
    embee.append(
        {
            "name": "Available at",
            "value": "```\n" + array_channels + "```",
            "inline": False,
        }
    )
    embee.append({"name": "Available", "value": str(data["available"]), "inline": True})
    embee.append(
        {"name": "Price", "value": str(data["fullPrice"]) + "€", "inline": False}
    )

    # based of the size of the line data["size"]

    line = len(data["size"].splitlines())
    number_of_element = line / 5
    number_of_element = int(number_of_element)

    for i in range(number_of_element):
        # create a list of 5 elements
        list_of_5 = data["size"].splitlines()[i * 5 : (i + 1) * 5]
        # join the list of 5 elements
        list_of_5 = "\n".join(list_of_5)

        embee.append({"name": "Size", "value": list_of_5, "inline": True})

    embee.append(
        {
            "name": "Resell Links",
            "value": f"[StockX]({stockx}) - [Klekt]({klekt}) - [Goat]({goat}) - [Restocks]({restock})",
            "inline": False,
        }
    )
    embee.append(
        {
            "name": "Useful Links",
            "value": f"[Cart]({cart}) - [App]({app})",
            "inline": False,
        }
    )
    current_time = datetime.now().strftime("%I:%M %p")

    data = {
        "username": "Nike Restock IT",
        "avatar_url": logo_uzumaki,
        "embeds": [
            {
                "title": data["title"],
                "url": url,
                "color": 12298642,
                "thumbnail": {"url": data["imageUrls"]},
                "fields": embee,
                "footer": {
                    "text": "Nike by Uzumaki" + " • Today at " + current_time,
                    "icon_url": logo_uzumaki,
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
