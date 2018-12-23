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
  and timestamp >= TIMESTAMP("{{from}}")
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
      and timestamp >= TIMESTAMP("{{from}}")
  )
order by
  time
;
