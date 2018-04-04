# hnanalysis
Hacker News analysis

# Usage

- Run `BigQuery/query.sql` on the BigQuery console.
- Save results as table, then export to google storage, finally save as `data/hn_from_2016.csv`.
- `make` to make go binary `hnanalysis`.
- `./hnanalysis data/hn_from_2016.csv results/result.csv` to analyse it.
