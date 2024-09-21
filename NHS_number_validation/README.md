# NHS Number Validation And Generation

This program validates a given NHS number and Generate a new valid one.

## NHS

The NHS NUMBER, the primary identifier of a PERSON, is a unique identifier for a PATIENT within the NHS in England and Wales.

The NHS NUMBER is 10 numeric digits in length.

## Usage

```
go build cmd

./cmd/main --operation=generate

./cmd/main --operation=validate --num={NHS_NUMBER}
```