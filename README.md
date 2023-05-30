### Sneaker Monitor
Monitoring sneakers store endpoint API to scrape new products and sends a discord webhook whenever a new limited product has been loaded or restocked.

### Supported Sites
The following sneaker stores are currently supported by the Sneaker Monitor:

1. Awlab (bypassing Cloudflare protection) 
2. Wethenew (consignment)
3. Wethenew (released)
4. Dover Street Market
5. Oqium
6. JD Sports
7. Nike
8. Nike-Restock
9. Vans
10. Ticket-Master

### Installation and Setup
To use the Sneaker Monitor, follow these steps:

Clone the repository to your local machine.\
Install the required dependencies by running pip install -r requirements.txt in the project directory.\
Configure the config.py file with your desired settings.(When required)\
Otherwise, you can run the main.py file with the default settings.(Filling up the webhook url directly in the file)\
Run the main.py file.

### Docker
To use the Sneaker Monitor with Docker, follow these steps:

Clone the repository to your local machine.
Build the Docker image by running docker build -t sneaker-monitor .
Run the Docker container by running docker run -d sneaker-monitor

### Contributing
If you would like to contribute to the Sneaker Monitor, please follow these steps:

Fork the repository.
Create a new branch for your changes.
Commit your changes to the new branch.
Create a pull request.

### Contact
If you have any questions, feel free to contact me on Discord at @Manuciao | YΞ#5388
