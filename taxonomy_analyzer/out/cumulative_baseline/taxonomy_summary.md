# Helm Error Taxonomy Report

Generated at: `2026-06-01 23:36:40 UTC`

Source catalog: `C:\Users\miabs\GolandProjects\helm-tests\helm_fetcher\catalog_cumulative.json`

## Totals

| Metric | Value |
|---|---:|
| Repositories | 963 |
| Helm runs | 9630 |
| Template failures | 2396 |
| Dependency failures | 277 |
| Classified errors | 2454 |
| Unclassified errors | 219 |

## Taxonomy by Kind

| Kind | Count |
|---|---:|
| `template` | 2197 |
| `dependency` | 257 |
| `unknown` | 219 |

## Taxonomy by SubKind

| SubKind | Count |
|---|---:|
| `template.nil_pointer` | 982 |
| `template.required_value` | 611 |
| `unknown.unclassified` | 219 |
| `dependency.missing_repository` | 153 |
| `template.missing_template` | 144 |
| `template.kube_version_incompatible` | 103 |
| `template.runtime_eval` | 96 |
| `template.values_schema_validation` | 96 |
| `template.yaml_render` | 55 |
| `dependency.missing_subchart` | 48 |
| `template.library_chart_not_installable` | 45 |
| `template.parse_error` | 30 |
| `template.type_mismatch` | 28 |
| `dependency.lock_file_out_of_sync` | 14 |
| `dependency.chart_validation` | 9 |
| `dependency.rate_limit` | 7 |
| `dependency.unpack_error` | 7 |
| `dependency.repo_update` | 6 |
| `dependency.version_resolution` | 6 |
| `template.invalid_value` | 6 |
| `dependency.cache_index_missing` | 4 |
| `dependency.network_dns` | 3 |
| `template.values_merge_error` | 1 |

## Unclassified Samples

- `collabnix/kubelabs` `D:\helm_clones_github\collabnix__kubelabs\Helm101\Wordpress` `dependency`: Error: cannot load Chart.yaml: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go struct field Metadata.deprecated of type bool
- `containers/kubernetes-mcp-server` `D:\helm_clones_github\containers__kubernetes-mcp-server\charts\kubernetes-mcp-server` `template`: Error: execution error at (kubernetes-mcp-server/templates/ingress.yaml:2:13): Ingress hostname must be specified  Use --debug flag to render out invalid YAML
- `WeBankFinTech/Prophecis` `D:\helm_clones_github\WeBankFinTech__Prophecis\install\Prophecis` `dependency`: Error: cannot load values.yaml: error reading yaml document: invalid Yaml document separator: --END RSA PRIVATE KEY-----"
- `IBM/charts` `D:\helm_clones_github\IBM__charts\community\aqua-enforcer` `template`: Error: execution error at (aqua-enforcer/templates/enforcer-token-secret.yaml:14:13): A valid .Values.enforcerToken entry required!  Use --debug flag to render out invalid YAML
- `IBM/charts` `D:\helm_clones_github\IBM__charts\community\aqua-scanner` `template`: Error: execution error at (aqua-scanner/templates/scanner-deployment.yaml:31:14): Please specify a username associated with the Scanner role!  Use --debug flag to render out invali...
- `Hydrospheredata/hydro-serving` `D:\helm_clones_github\Hydrospheredata__hydro-serving\helm` `dependency`: Error: cannot load Chart.yaml: error unmarshaling JSON: while decoding JSON: json: cannot unmarshal string into Go value of type v2.Metadata
- `lianqingsec/NucleiPocGather` `D:\helm_clones_github\lianqingsec__NucleiPocGather\poc\other` `dependency`: Error: chart file "signatures_1.yaml" is larger than the maximum file size 5242880
- `IndustryFusion/DigitalTwin` `D:\helm_clones_github\IndustryFusion__DigitalTwin\helm\charts\pgrest` `template`: Error: execution error at (pgrest/templates/ingress.yaml:3:4): ingressType must either be "traefik" or "nginx"  Use --debug flag to render out invalid YAML
- `acuvity/mcp-servers-registry` `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-21st-dev-magic\charts\mcp-server-21st-dev-magic` `template`: Error: execution error at (mcp-server-21st-dev-magic/templates/secrets.yaml:10:9): required value for secrets.TWENTY_FIRST_API_KEY either as .value or .valueFrom.name and .valueFro...
- `acuvity/mcp-servers-registry` `D:\helm_clones_github\acuvity__mcp-servers-registry\mcp-server-adfin\charts\mcp-server-adfin` `template`: Error: execution error at (mcp-server-adfin/templates/secrets.yaml:10:9): required value for secrets.ADFIN_PASSWORD either as .value or .valueFrom.name and .valueFrom.key  Use --de...
