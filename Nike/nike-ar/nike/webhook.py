from datetime import datetime
import requests


def webhook(productName, linkText, sku, image, availableQuantity, Price, talle):
    print("[WEBHOOK] Sending webhook...")
    logo = "https://i.imgur.com/4M34hi2.png"

    stockx = f"https://stockx.com/search?s={sku}"
    klekt = f"https://www.klekt.com/brands?search={sku}"
    goat = f"https://www.goat.com/sneakers?query={sku}"
    restock = f"https://restocks.net/en/shop/?q={sku}"
    cart = "https://www.nike.com.ar/checkout/#/cart"

    webhook = ""
    url = f"https://www.nike.com.ar{linkText}"

    embee = []

    array_size = ""
    for size in talle:
        array_size += f"{size}\n"

    embee.append({"name": "SKU", "value": sku, "inline": True})
    embee.append({"name": "Region", "value": ":flag_ar:", "inline": False})
    embee.append({"name": "Quantity", "value": str(availableQuantity), "inline": True})
    embee.append({"name": "Price", "value": str(Price) + "$", "inline": False})

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
        "username": "Nike AR",
        "avatar_url": logo,
        "embeds": [
            {
                "title": productName,
                "url": url,
                "color": 12058624,
                "thumbnail": {"url": image},
                "fields": embee,
                "footer": {
                    "text": "Nike AR" + " • Today at " + current_time,
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
