from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.common.keys import Keys
import time

driverpath = '/usr/local/bin/chromedriver'
driver = webdriver.Chrome()

# driver.get("https://google.com")
driver.get("https://techwithtim.net")

search = driver.find_element(By.ID, "s")
search.send_keys("test")
search.send_keys(Keys.RETURN)

# Optionally, wait for the page to load dynamically
time.sleep(5)

# Close the driver after use
driver.quit()
