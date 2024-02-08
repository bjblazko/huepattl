# Hüpattl!

## Requirements

### BigQuery

- enable API
- table `site.blogs` with schema:
 
| Field name | Type           | Mode       |
|:-----------|:---------------|:-----------|
| date       | `DATETIME`     | `REQUIRED` |
| author     | `STRING(255)`  | `REQUIRED` |
| title      | `STRING(255)`  | `REQUIRED` |
| preview    | `STRING(1000)` | `REQUIRED` |
| bucket     | `STRING(255)`  | `REQUIRED` |
| filename   | `STRING(255)`  | `REQUIRED` |
