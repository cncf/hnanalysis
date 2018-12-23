# hnanalysis
Hacker News analysis

# Usage

- Execute script [get_bigquery_data.sh](https://github.com/cncf/hnanalysis/blob/master/util_sh/get_bigquery_data.sh), for example: `./util_sh/get_bigquery_data.sh 'YYYY-MM-DD HH:MI:SS'`.
- Or run [BigQuery/query.sql](https://github.com/cncf/hnanalysis/blob/master/BigQuery/query.sql) on the BigQuery console replacing `{{from}}` with `YYYY-MM-DD HH:MI:SS`.
- If run on BigQuery console then save results as table, then export to google storage, finally save as [data/hn.csv](https://github.com/cncf/hnanalysis/blob/master/data/hn.csv) (`get_bigquery_data.sh` does it automatically).
- Update [jobs.yaml](https://github.com/cncf/hnanalysis/blob/master/jobs.yaml) to specify job postings to search for.
- Run [hnanalysis.sh](https://github.com/cncf/hnanalysis/blob/master/hnanalysis.sh) to generate [results/result.csv](https://github.com/cncf/hnanalysis/blob/master/results/result.csv).
- Import the final [result.csv](https://github.com/cncf/hnanalysis/blob/master/results/result.csv) into [this](https://docs.google.com/spreadsheets/d/1nVTk7rA9zObe0BgkWrCj34WdmfGllQEbePnGpO0L5Ys/edit?usp=sharing) Google Sheet (B1 cell).
- Final chart is [here](https://docs.google.com/spreadsheets/d/1nVTk7rA9zObe0BgkWrCj34WdmfGllQEbePnGpO0L5Ys/edit#gid=1756635924).
