import requests
import json

from bs4 import BeautifulSoup, Comment, Tag
from tinyhtml import html, h, frag

auth = 'Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsIng1dCI6InEtMjNmYWxldlpoaEQzaG05Q1Fia1A1TVF5VSIsImtpZCI6InEtMjNmYWxldlpoaEQzaG05Q1Fia1A1TVF5VSJ9.eyJhdWQiOiI1YjIwN2M1Zi04YWJmLTQ2NDUtYjM0MC1hMTc4YzcxZWU1N2IiLCJpc3MiOiJodHRwczovL3N0cy53aW5kb3dzLm5ldC9mY2IyYjM3Yi01ZGEwLTQ2NmItOWI4My0wMDE0YjY3YTdjNzgvIiwiaWF0IjoxNzEyMjg0NTQ2LCJuYmYiOjE3MTIyODQ1NDYsImV4cCI6MTcxMjI4ODQ0NiwiYWlvIjoiRTJOZ1lGaVZFVm14NEpKMDFOV2lsMEZ2RFRxaUFRPT0iLCJhcHBpZCI6IjViMjA3YzVmLThhYmYtNDY0NS1iMzQwLWExNzhjNzFlZTU3YiIsImFwcGlkYWNyIjoiMSIsImlkcCI6Imh0dHBzOi8vc3RzLndpbmRvd3MubmV0L2ZjYjJiMzdiLTVkYTAtNDY2Yi05YjgzLTAwMTRiNjdhN2M3OC8iLCJvaWQiOiI3ZjA1MTBlZS1kNjhiLTRmOWQtYjhlMi1iNzVhMDIzODgyMmMiLCJyaCI6IjAuQVFzQWU3T3lfS0JkYTBhYmd3QVV0bnA4ZUY5OElGdV9pa1ZHczBDaGVNY2U1WHNMQUFBLiIsInN1YiI6IjdmMDUxMGVlLWQ2OGItNGY5ZC1iOGUyLWI3NWEwMjM4ODIyYyIsInRpZCI6ImZjYjJiMzdiLTVkYTAtNDY2Yi05YjgzLTAwMTRiNjdhN2M3OCIsInV0aSI6IlFHTHdQY2xYSTAyMmlGdjhUblFoQUEiLCJ2ZXIiOiIxLjAiLCJlbWFpbF92ZXJpZmllZCI6InRydWUiLCJodHRwczovL2JheWVyLmNvbS9hcHBfZGlzcGxheW5hbWUiOiJIYXlzdGFjay1NQVBTLVVJLU5QIn0.Q9qovioE4ryjQKqjXDZXXixGr2ouxexEPIlhEFXdzoP7SQo5yRt5U-SZegA1U3-AtZUKPOe_Hr-KgCsqN3mn47W9LsVvnKEhQrdjI_sGmDqLtKAFv4TIkROYgGWnNO5vXktNE7IWIJh-Mly31ensU2tBSfzlgV_PoHJMuNxllY5Arf1jBRzU47LWHrH3MLJrkh52QhdAYXKcPZ_nk8m3EhAR4aY2XjuHQVntMqEvCEkIho-mB5PM2yc1loYrisRavGWC2AyjS_GKTigxohC45O_2brOrHMcKxGX_u6X1RBnBECwdxGVYNlYRfb8G0vzgz1FKHeGQXqwaeVcKuFPOwg'

def call_haystack():
    url = 'https://haystack.bayer.com/api.php'

    headers = {
        'content-type': 'application/json', 
        'Authorization': auth
    }

    payload = {
        'action': 'parse',
        'format': 'json',
        'prop': 'wikitext|text', # get wikitext and text
        'page': 'Above_Allocation'
    }

    r = requests.get(
        url=url,
        params=payload,
        headers=headers,
        verify=False
    )

    return r

def get_results(r):
    t = json.loads(r.text)

    ty = type(t)

    # print(t)
    # print(ty)

    t2 = t['parse']['text']['*']
    wi = t['parse']['wikitext']['*']

    j = r.json()['parse']['text']['*']

    # with open("text.txt", "w") as f:
    #     f.write(json.dumps(t))

    # with open("text.json", "w") as f:
    #     f.write(json.dumps(t))

    # with open("text.html", "w") as f:
    #     f.write(json.dumps(t2))

    # with open("wiki.html", "w") as f:
    #     f.write(json.dumps(wi))

    with open("js.json.html", "w") as f:
        f.write(json.dumps(j))

    return j

def get_soup(j):
    soup = BeautifulSoup(j, 'html.parser')

    for element in soup(text=lambda text: isinstance(text, Comment)):
        element.extract()

    # print(soup.prettify())
        
    # with open("res.html", "w") as f:
    #     f.write(soup.prettify())
    return soup
    
def get_updates(soup):
    tables = soup.find_all('table')
    # # print(tables)
    # print(type(tables))
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
                # print(row)
                # print('----------------')
                if not row.find_all('th'):
                    data = [el.text.strip() for el in row.find_all('td')]
                    print(data)
                    rows.append(data)
            print(rows)
            return rows

def get_authorized_date(soup):
    divs = soup.find_all('div')
    # # print(tables)
    # print(type(tables))
    print(len(divs))
        
    final_auth_value = []
    for div in divs:
        # tr, th, td 
        if "metadata-content" == div['class'][0] and "properties-content" == div['class'][1] and div.find('h2').find('span')['id'] == 'Properties':
            # now find the auth data info
            for div2 in div:
                # print(type(div2))
                for div3 in div2:
                    if isinstance(div3, Tag) and 'class' in div3.attrs and "metadata-row" == div3['class'][0]:
                        # print(div3)
                        p = div3.find_all('p')
                        if 'Additional Reference Documentation' in p[0]:

                            lines = str(p[1]).split('<br/>')
                            index = 0
                            for i, line in enumerate(lines):
                                if 'Authorized' in line:
                                    # print(line)
                                    index = i
                                    break
                            thing = lines[index+1]
                            final_auth_value.append(thing)
                            # print('-----------------------')
                        #     print(div3)
                        # print('-----------------------')

    if(len(final_auth_value) == 1):
        return final_auth_value[0]

def gen_html_table(updates, auth_date):
    
    headers = ['Authorized By', 'Date']

    def create_header():
        open = "<th>"
        close = "</th>"
        return''.join([open + h + close for h in headers])

    def create_data(r):
        open = "<td>"
        close = "</td>"
        return ''.join([open + d + close for d in r])

    def create_rows():
        h_open = '<tr bgcolor="#F5DEB3" style="height: 2em;">'
        d_open = '<tr>'
        close = "</tr>"
        header = h_open + create_header() + close
        # TODO: handle no udpates or auth date
        dataauthrow = '' if not auth_date else d_open + create_data(auth_date.split('on ')) + close
        dataupdates = '' if not updates else ''.join([d_open + create_data(r) + close for r in updates])
        return header + dataauthrow + dataupdates
    
    def create_table():
        open = '<table style="border: 1px solid #F5DEB3; width: 70%;">'
        close = "</table>"
        return open + create_rows() + close

    def create_html():
        open = """'<!DOCTYPE html>
        <html lang="en">"""
        close = "</html>"
        return open + create_table() + close
        

    html_str = create_html()
    soup_str = BeautifulSoup(html_str, 'html.parser')
    print(soup_str.prettify())

    # table = Tag(new_soup, "table")
    # tr = Tag(new_soup, "tr")
    # new_soup.append(html)
    # html.append(table)
    # table.append(tr)
    # for attr in headers:
    #     th = Tag(new_soup, "th")
    #     tr.append(th)
    #     th.append(attr)
    
    # print(new_soup.prettify())

r = call_haystack()
j = get_results(r)
soup = get_soup(j)
updates = get_updates(soup)
auth_date = get_authorized_date(soup)

gen_html_table(updates, auth_date)
print(updates)
print(auth_date)

print('done')