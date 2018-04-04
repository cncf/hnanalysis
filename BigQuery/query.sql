#standardSQL
select
  id,
  parent,
  time,
  timestamp,
  type,
  ranking,
  descendants,
  title,
  text 
from
  `bigquery-public-data.hacker_news.full`
where
  REGEXP_CONTAINS(title, r'^Ask HN: Who is hiring\?\s+\(\w+\s+\d+\)$')
  and timestamp >= TIMESTAMP("2016-01-01 00:00:00")
union all select
  id,
  parent,
  time,
  timestamp,
  type,
  ranking,
  descendants,
  title,
  text 
from
  `bigquery-public-data.hacker_news.full`
where
  parent in (
    select
      id
    from
      `bigquery-public-data.hacker_news.full`
     where
      REGEXP_CONTAINS(title, r'^Ask HN: Who is hiring\?\s+\(\w+\s+\d+\)$')
      and timestamp >= TIMESTAMP("2016-01-01 00:00:00")
  )
order by
  time
;
