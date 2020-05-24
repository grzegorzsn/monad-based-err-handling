def main():
    try:
        run_everything()
    except Exception as e:
        print("Runtime err:", e)


def run_everything():
    file_data = fetch_data_from_file()
    db_data = fetch_data_from_db()
    processed_data = processData(file_data, db_data)
    save_processed_data(processed_data)


def fetch_data_from_file():
    with open("some_filename") as f:
        raw_bytes = f.read()
    data = parse_data(raw_bytes)
    validate(data)
    return data


# -----------------------------------------------
# Code below is not implemented. You can skip it.
# -----------------------------------------------


def validate(data):
    raise NotImplementedError()


def parse_data(raw_bytes):
    raise NotImplementedError()


def fetch_data_from_db():
    raise NotImplementedError()


def save_processed_data(processed_data):
    raise NotImplementedError()


def processData(file_data, db_data):
    raise NotImplementedError()


if __name__ == "__main__":
    main()
