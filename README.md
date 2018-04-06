# hnanalysis
Hacker News analysis

# Usage

- Run [BigQuery/query.sql](https://github.com/cncf/hnanalysis/blob/master/BigQuery/query.sql) on the BigQuery console.
- Save results as table, then export to google storage, finally save as [data/hn.csv](https://github.com/cncf/hnanalysis/blob/master/data/hn.csv).
- Update [jobs.yaml](https://github.com/cncf/hnanalysis/blob/master/jobs.yaml) to specify job postings to search for.
- Run [hnanalysis.sh](https://github.com/cncf/hnanalysis/blob/master/hnanalysis.sh) to generate [results/result.csv](https://github.com/cncf/hnanalysis/blob/master/results/result.csv).
- Import the final [result.csv](https://github.com/cncf/hnanalysis/blob/master/results/result.csv) into [this](https://docs.google.com/spreadsheets/d/1nVTk7rA9zObe0BgkWrCj34WdmfGllQEbePnGpO0L5Ys/edit?usp=sharing) Google Sheet.
- Final chart is [here](https://docs.google.com/spreadsheets/d/1nVTk7rA9zObe0BgkWrCj34WdmfGllQEbePnGpO0L5Ys/edit#gid=1756635924).
