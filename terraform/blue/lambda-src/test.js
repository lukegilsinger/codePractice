const aws = require('aws-sdk');
const request = require('request');
const https = require('https')

key = 'YC2Z5JLI4LD1EPUN'

// replace the "demo" apikey below with your own key from https://www.alphavantage.co/support/#api-key
// var url = 'https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=IBM&interval=5min&apikey=demo';
var url = `https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=IBM&interval=5min&apikey=${key}`;

writeS3 = async (metadata, filename = 'data/test.json') => {
  const s3 = new aws.S3();
  const s3Bucket = 'du-blue-test-bucket'; // replace with your bucket name
  const objectName = filename; // File name which you want to put in s3 bucket
  const objectData = JSON.stringify(metadata); // file data you want to put
  const objectType = 'application/json'; // type of file

  console.log('writing ', filename)

  try {
    const params = {
       Bucket: s3Bucket,
       Key: objectName,
       Body: objectData,
       ContentType: objectType,
    };
    const result = await s3.putObject(params).promise();
    console.log(`File uploaded successfully at https:/` + s3Bucket +   `.s3.amazonaws.com/` + objectName);
    return result
  } catch (error) {
    console.log(error);
    return 1
  }
}

app = async () => {
  console.log('hello-world')

  console.log(await chuck())

  await writeS3({test: 1}, 'data/thing.json')

  console.log(url)

  // request.get({
    //   url: url,
    //   json: true,
    //   headers: {'User-Agent': 'request'}
    // }, (err, res, data) => {
    //   if (err) {
    //     console.log('Error:', err);
    //   } else if (res.statusCode !== 200) {
    //     console.log('Status:', res.statusCode);
    //   } else {
    //     console.log(data["Meta Data"])
    //     console.log(Object.keys(data))
    //     try {
    //       writeS3(data["Meta Data"])
    //     } catch (e) {
    //       console.log(e)
    //     }
    //   }
  // });
  return 0
}


async function chuck() {
  console.log('getting fact')
  const res = await fetch('https://api.chucknorris.io/jokes/random')
  const randomFact = JSON.parse(res).value

  return randomFact
}

async function fetch(url) {
  return new Promise((resolve, reject) => {
    const request = https.get(url, { timeout: 1000 }, (res) => {
      if (res.statusCode < 200 || res.statusCode > 299) {
        return reject(new Error(`HTTP status code ${res.statusCode}`))
      }

      const body = []
      res.on('data', (chunk) => body.push(chunk))
      res.on('end', () => {
        const resString = Buffer.concat(body).toString()
        resolve(resString)
      })
    })

    request.on('error', (err) => reject(err))
    request.on('timeout', (err) => {
      console.log('timed out', err)
      // reject(err)
    })
  })
}

exports.app = app

app()