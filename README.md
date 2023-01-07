# ROAgen

ROAgen is a program used to take the RIPE RPKI validator's export JSON and turn it into ROA files for bird2.

You can download the export JSON at https://rpki-validator.ripe.net/api/export.json.

## Usage

The following flags must be passed when running ROAgen:
- -data: Location of the RIPE RPKI validator export file.
- -out: Directory where ROAgen should save the ROA files. Directory must already exist

Two ROA files (`roa4.conf` and `roa6.conf`) will be placed in the output directory.

## Performance

ROAgen is fast, taking only a few seconds to generate the ROA files. It uses roughly 200MB of RAM.