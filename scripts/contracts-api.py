
import requests
import argparse
import json

def call_list_operational_accounts(account_ids, source_system_code, token):
    # Construct the URL with the account IDs
    account_id_param = '&'.join([f'account_id={account_id}' for account_id in account_ids])
    url = f'https://nonprod.customer360.bcs.bayer.com/api/v1/accountmanagement/operationalaccounts?source_system_code={source_system_code}&returnBiDirectionalRelatedAccounts=true&source_system_code=&{account_id_param}'

    # Set up the headers with the authorization token
    headers = {
        'Authorization': f'Bearer {token}'
    } 

    # Make the GET request
    response = requests.get(url, headers=headers, verify=False)

    if response.status_code == 200:
        return response.json()  # Return the JSON response if successful
    else:
        print(f'Error: {response.status_code} - {response.text}')
        return None

def get_related_ird_account(account):
    relatedAccounts = account['relatedAccounts']
    # print(relatedAccounts)
    relOplAccts = relatedAccounts['operationalAccounts']
    # print(relOplAccts)
    irdAccount = {}
    for a in relOplAccts:
        # print(a)
        if a['sourceSystemCode'] == 'IRD-US':
            irdAccount = a
    return irdAccount

def get_contracts(account):
    contracts = []
    for c in account['contracts']:
        if c['contractType'] == 'MBSD':
            contracts.append(c)
    
    return contracts

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='Get operational accounts from the API.')
    parser.add_argument('--account_ids', nargs='+', required=True, help='List of account IDs to query')
    parser.add_argument('--token', required=True, help='Authorization token')

    args = parser.parse_args()

    # print("ARGS:", args)

    # Call the function with provided account IDs and token
    accounts = call_list_operational_accounts(args.account_ids, 'SAP-QBC-C', args.token)
    if accounts:
        # account = json.loads(accounts)
        operationalaccounts = accounts['operationalaccounts']
        # print(operationalaccounts)
        ird_account = get_related_ird_account(operationalaccounts[0])
        print(ird_account)
        if ird_account:
            code = ird_account['sourceSystemCode']
            id = ird_account['sourceId']
            res2 = call_list_operational_accounts([id], code, args.token)
            print(res2)
            contracts = get_contracts(res2['operationalaccounts'][0])
            print(json.dumps(contracts, indent=4))