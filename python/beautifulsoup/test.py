from bs4 import BeautifulSoup

# Opening the html file 
HTMLFile = open("temp.html", "r")
# Reading the file 
index = HTMLFile.read() 

soup = BeautifulSoup(index, 'html.parser')
# print(soup.prettify()) 

#write file
with open("res.html", "w") as f:
    f.write(soup.prettify())

tables = soup.find_all('table')
# print(tables)
print(type(tables))
print(len(tables))
    
for table in tables:
    # tr, th, td 
    h = table.find('tr') #first row
    header = [element.text.strip() for element in h.find_all('th')]

    print(header)

    if len(header) == 2 and 'Authorized By' in header[0] and 'Date' in header[1]:
        print('Found UPDATES table')
        rows = []
        for i, row in enumerate(table.find_all('tr')):
            if not row.find_all('th'):
                data = [el.text.strip() for el in row.find_all('td')]
                # print(data)
                rows.append([data])
        # print(rows)
    


divs = soup.find_all('div')
print(type(divs))
print(len(divs))

for div in divs:
    # print(div['class'])
    if "metadata-content" in div['class'][0]:
        # print(type(div.contents))
        for c in div.contents:
            print(type(c))
            print(c)
        # for c in div.children:
        #     print(type(c))
        #     print(c)
        #     print('----------------------------')
        # for d in div:
        #     print(d)

print('DONE')