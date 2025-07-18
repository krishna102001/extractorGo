## Extractor Go

Extractor go is a fast & secure backend services to extract the embedded images from the pdf or convert the each pdf pages into a images. Its make the password protected zip files.

# Features

- Extract the embedded images
- Convert the pdf file into images
- password protected zip files
- support base64 or form file

# How its works

- To extract images from pdf make `POST` request to `api/v1/extract-pdf-image`.

  ```bash request-body(json)
       {"input_pdf_file":"base64 text"}
  ```

  OR

  ```bash request-body(form-data)
       file:"sample.pdf"
  ```

  ```bash response
  { "data":{
    "extract_id":"uuid",
    "response_url":"https://cloudinary.com/......"
    }
  }
  ```

- To get the saved response back make `GET` request to `api/v1/extract/:id`

  ```bash response
  { "data":{
    "extract_id":"uuid",
    "response_url":"https://cloudinary.com/......"
    }
  }
  ```

- To convert pdf pages into images make `POST` request to `api/v1/convert-pdf-image`.

  ```bash request-body(json)
       {"input_pdf_file":"base64 text"}
  ```

  OR

  ```bash request-body(form-data)
       file:"sample.pdf"
  ```

  ```bash response
  { "data":{
    "convert_id":"uuid",
    "response_url":"https://cloudinary.com/......"
    }
  }
  ```

- To get the saved response back make `GET` request to `api/v1/convert/:id`
  ```bash response
  { "data":{
    "convert_id":"uuid",
    "response_url":"https://cloudinary.com/......"
    }
  }
  ```
