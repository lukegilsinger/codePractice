print('hello')

l1 = [
    "3787c957-ff20-4d7e-9649-21087404369b",
    "3493ce49-382b-4a05-86b9-ae4c38aa99e7",
    "3778119a-0970-447e-b389-3ec6c990c095",
    "8da3998e-ca97-485d-a554-1f5eb505465e",
    "52d33351-2376-42b1-8d5f-5da1a9c07ef5",
    "14c1383b-b625-468f-ac80-670e77886e88",
    "0bda1398-5216-4979-bc08-3c2e9d8d2792",
    "18e05118-f730-4c15-9c7c-f152c033734f",
    "26544301-48d6-4bc1-b72a-c02bb104195d",
    "fdf2626c-c488-4e54-86d9-ac9823f04a13",
    "e6f85bf9-00eb-40ce-800f-158239ffa7d8",
    "1105d8a8-46d5-4003-bae4-a11771d4f87e",
    "3ca4bdee-7e72-4bd2-b93d-305c31ec1872",
    "848c3d85-f286-478a-8fe8-b7c2aff33805",
    "e5a4fdf9-6f75-4dac-929a-3d200bc50696",
    "904e4b09-7134-4881-b4d2-74bfd63840e4",
    "b91a2c6d-40c2-4e8e-8f57-d2cf1d917be6",
    "bdb3293d-25a3-412b-a4ea-f90873ff000b",
    "fe212e6f-c941-408f-b6da-e805be02b8ad",
    "ddfd6cf3-e609-47c5-8e22-517882298593",
    "3ccc5ea9-c79e-44a9-b6a6-45392eccee22"
]

l2 = [
    "3787c957-ff20-4d7e-9649-21087404369b",
    "3493ce49-382b-4a05-86b9-ae4c38aa99e7",
    "1c377708-df71-4f2f-9441-5be882b03f27",
    "b99575eb-5fa5-4160-9d8f-50db0cc8549c",
    "8c96be4c-c0eb-4b25-8b5f-24c5553b0598",
    "ed245d11-910d-4db0-bf2d-b01907898a69",
    "4b2c38d4-36f2-42f3-ba26-f6538c07a0f7",
    "494f49e4-f646-44e2-b617-db96d1140551",
    "bc157681-a329-4318-b5f7-325200457899",
    "b581e5cd-e244-4a87-b516-07e874d4f16f",
    "9b904728-cb21-414a-b99b-01c3d32ddcf3",
    "590be5f3-99a7-43d2-a768-01acd7cf9839",
    "34987faa-c4b0-40df-8dd9-e96f79d0247b",
    "26544301-48d6-4bc1-b72a-c02bb104195d",
    "e14242ff-a546-418d-898c-bb17397abd38",
    "081db232-234c-4de0-8287-aa8f10d454c5",
    "ff0e42d5-6287-4bf9-b970-2fa1796132e9"
]

unique = list(set(l2))
print(unique)

lx = l1 + l2

unique = list(set(lx))
print(len(l1))
print(len(l2))

print(unique)

for i in unique:
    print(i)