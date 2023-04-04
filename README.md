### Sneaker Monitor
The Sneaker Monitor is a Python-based tool designed to monitor sneaker stores' endpoint APIs and scrape new products, then send notifications via Discord webhook whenever a new limited product has been loaded or restocked.

### Supported Sites
The following sneaker stores are currently supported by the Sneaker Monitor:

1. Awlab (bypassing Cloudflare protection) 
2. Wethenew (consignment)
3. Wethenew (released)
4. Dover Street Market
5. Oqium
6. JD Sports

### Installation and Setup
To use the Sneaker Monitor, follow these steps:

Clone the repository to your local machine.
Install the required dependencies by running pip install -r requirements.txt in the project directory.
Configure the config.py file with your desired settings.
Run the main.py file.

### Docker
To use the Sneaker Monitor with Docker, follow these steps:

Clone the repository to your local machine.
Build the Docker image by running docker build -t sneaker-monitor .
Run the Docker container by running docker run -d sneaker-monitor

### Configuration
The Sneaker Monitor can be configured by editing the config.py file. The following settings are available:

DISCORD_WEBHOOK_URL: The Discord webhook URL to send notifications to.
DISCORD_WEBHOOK_AVATAR_URL: The Discord webhook avatar URL to use for notifications.
DISCORD_WEBHOOK_USERNAME: The Discord webhook username to use for notifications.
DISCORD_WEBHOOK_EMBED_COLOR: The Discord webhook embed color to use for notifications.
DISCORD_WEBHOOK_EMBED_IMAGE: The Discord webhook embed image to use for notifications.

### Contributing
If you would like to contribute to the Sneaker Monitor, please follow these steps:

Fork the repository.
Create a new branch for your changes.
Commit your changes to the new branch.
Create a pull request.

### Contact
If you have any questions, feel free to contact me on Discord at @Manuciao | YΞ#5388