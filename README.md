# Zuri NFT parser

> An Nft parser for Zuri NFT's

## Getting started

# Requirements

- Clone the git repository
- move your teams csv file or the general csv file to the root folder of the repository
- make sure there are no spaces in your file name eg. TeamBevel.csv
- Open the terminal and go into the repositiory directory
- run the following command

```sh
./main --input {{insert path/to/csv file here}}
```

### Result

A new folder will be created with the json files of the nfts. An output.csv file will also be created that contains the sha265 hash of the json files.

### Example json file

```json
{
  "format": "CHIP-0007",
  "name": "adewale-the-amebo",
  "description": "Adewale always wants to be in everyone's business.",
  "minting_tool": "TeamBevel",
  "sensitive_content": false,
  "series_number": 1,
  "series_total": 420,
  "collection": {
    "name": "Zuri NFT Tickets for Free Lunch",
    "id": "b774f676-c1d5-422e-beed-00ef5510c64d",
    "attributes": [
      {
        "type": "description",
        "value": "Rewards for accomplishments during HNGi9."
      }
    ]
  },
  "gender": "",
  "uuid": "cad316c3-37f8-4b27-9f53-9d803bfcfee7",
  "hash": "E71B82700FFE4103D1AD92D1F4D92598945A480ABD0526D0110381362F8B3FAC"
}
```
### Example csv output file

```csv
    Series Number,Filename,Description,Gender,UUID,Hash,sha256
    1,adewale-the-amebo,Adewale always wants to be in everyone's business.,male,cad316c3-37f8-4b27-9f53-9d803bfcfee7,E71B82700FFE4103D1AD92D1F4D92598945A480ABD0526D0110381362F8B3FAC,efd8bffe297174b8c9650f98268aa0451bfe13d8176702ea9ea1db0f81ea2485
    2,alli-the-queeny,Alli is an LGBT Stan.,male,a57ecda4-a26f-4e97-a757-e1800af39aab,E1ECE1FFACC43C3A856DE6D02185DC046518EE91D7AB934D9AD4051A698046E2,3e0a4bcc49dc3543451c08cf24ad13b14633dfa5f5d4f3f3f502c8a7944bcbf7
```

Please feel free to reach out on Slack @Richdotcom for enquries or clarifications or if you would like to update fields or data
