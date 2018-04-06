# hnanalysis
Hacker News analysis

# Usage

- Run [BigQuery/query.sql](https://github.com/cncf/hnanalysis/blob/master/BigQuery/query.sql) on the BigQuery console.
- Save results as table, then export to google storage, finally save as [data/hn.csv](https://github.com/cncf/hnanalysis/blob/master/data/hn.csv).
- Update [jobs.yaml](https://github.com/cncf/hnanalysis/blob/master/jobs.yaml) to specify job postings to search for.
- Run [hnanalysis.sh](https://github.com/cncf/hnanalysis/blob/master/hnanalysis.sh) to generate [results/result.csv](https://github.com/cncf/hnanalysis/blob/master/results/result.csv).
