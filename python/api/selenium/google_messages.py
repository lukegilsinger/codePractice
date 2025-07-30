from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.common.keys import Keys
import time

driverpath = '/usr/local/bin/chromedriver'
driver = webdriver.Chrome()

# driver.get("https://google.com")
driver.get("https://messages.google.com/web")

# Give time for manual login or QR code scan
input("Please log in manually and press Enter to continue...")

# Optionally, wait for the page to load dynamically
time.sleep(5)

# Example: find message bubbles (adjust the selector as necessary)
messages = driver.find_elements(By.CSS_SELECTOR, ".text-message")  # CSS selector will vary

# Loop through and print message content
for message in messages:
    print(message.text)


# Optionally, wait for the page to load dynamically
time.sleep(5)

# Close the driver after use
driver.quit()
