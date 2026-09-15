module github.com/kleister/kleister-api

go 1.26.5

require (
	github.com/Machiel/slugify v1.0.1
	github.com/aws/aws-sdk-go-v2 v1.47.0
	github.com/aws/aws-sdk-go-v2/config v1.33.5
	github.com/aws/aws-sdk-go-v2/credentials v1.20.5
	github.com/aws/aws-sdk-go-v2/service/s3 v1.113.1
	github.com/aws/smithy-go v1.28.1
	github.com/cenkalti/backoff/v7 v7.0.0
	github.com/coreos/go-oidc/v3 v3.21.0
	github.com/dchest/authcookie v0.0.0-20190824115100-f900d2294c8e
	github.com/dchest/uniuri v1.2.0
	github.com/disintegration/imaging v1.6.2
	github.com/drexedam/gravatar v0.0.0-20210327211422-e94eea8c338e
	github.com/drone/funcmap v0.0.0-20240227160611-7e19e9cd5a1c
	github.com/elimity-com/scim v0.0.0-20260728105928-2641426a1539
	github.com/gabriel-vasile/mimetype v1.4.15
	github.com/getkin/kin-openapi v0.149.0
	github.com/go-chi/chi/v5 v5.3.2
	github.com/go-chi/render v1.0.3
	github.com/go-openapi/runtime/server-middleware v0.33.2
	github.com/go-ozzo/ozzo-validation/v4 v4.4.1
	github.com/go-sql-driver/mysql v1.10.1
	github.com/gobwas/glob v1.0.0
	github.com/golang-jwt/jwt/v5 v5.3.1
	github.com/invopop/jsonschema v0.14.0
	github.com/joho/godotenv v1.5.1
	github.com/kleister/go-forge v1.0.1
	github.com/kleister/go-minecraft v1.11.1
	github.com/mcuadros/go-version v0.0.0-20190830083331-035f6764e8d2
	github.com/oapi-codegen/nethttp-middleware v1.2.0
	github.com/oapi-codegen/runtime v1.7.0
	github.com/oklog/run v1.2.0
	github.com/prometheus/client_golang v1.24.1
	github.com/rrivera/identicon v0.0.0-20240116195454-d5ba35832c0d
	github.com/rs/zerolog v1.35.1
	github.com/scim2/filter-parser/v2 v2.3.1
	github.com/spf13/cobra v1.10.2
	github.com/spf13/viper v1.21.0
	github.com/stretchr/testify v1.12.1
	github.com/uptrace/bun v1.2.18
	github.com/uptrace/bun/dialect/mysqldialect v1.2.18
	github.com/uptrace/bun/dialect/pgdialect v1.2.18
	github.com/uptrace/bun/dialect/sqlitedialect v1.2.18
	github.com/uptrace/bun/driver/pgdriver v1.2.18
	github.com/uptrace/bun/driver/sqliteshim v1.2.18
	github.com/vincent-petithory/dataurl v1.0.0
	go.uber.org/mock v0.6.0
	golang.org/x/crypto v0.57.0
	golang.org/x/oauth2 v0.37.0
)

require (
	filippo.io/edwards25519 v1.2.0 // indirect
	github.com/ajg/form v1.5.1 // indirect
	github.com/apapsch/go-jsonmerge/v2 v2.0.0 // indirect
	github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.7.20 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.20.0 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.5.3 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.8.3 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.5.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.13.19 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.11.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.14.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.20.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/signin v1.10.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.38.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.43.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.51.0 // indirect
	github.com/bahlo/generic-list-go v0.2.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/buger/jsonparser v1.2.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/dlclark/regexp2 v1.12.0 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/fsnotify/fsnotify v1.9.0 // indirect
	github.com/go-jose/go-jose/v4 v4.1.4 // indirect
	github.com/go-openapi/jsonpointer v1.0.0 // indirect
	github.com/go-viper/mapstructure/v2 v2.5.0 // indirect
	github.com/google/pprof v0.0.0-20260709232956-b9395ee17fa0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/mattn/go-colorable v0.1.15 // indirect
	github.com/mattn/go-isatty v0.0.22 // indirect
	github.com/mattn/go-sqlite3 v1.14.34 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/ncruces/go-strftime v1.0.0 // indirect
	github.com/oasdiff/yaml v0.1.1 // indirect
	github.com/oasdiff/yaml3 v0.0.14 // indirect
	github.com/pb33f/ordered-map/v2 v2.3.1 // indirect
	github.com/pelletier/go-toml/v2 v2.4.3 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_model v0.6.2 // indirect
	github.com/prometheus/common v0.70.1 // indirect
	github.com/prometheus/procfs v0.21.1 // indirect
	github.com/puzpuzpuz/xsync/v3 v3.5.1 // indirect
	github.com/q-uint/parser v0.3.1 // indirect
	github.com/q-uint/xsd-datetime v1.0.0 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	github.com/rs/xid v1.6.0 // indirect
	github.com/sagikazarmark/locafero v0.11.0 // indirect
	github.com/santhosh-tekuri/jsonschema/v6 v6.0.3 // indirect
	github.com/sourcegraph/conc v0.3.1-0.20240121214520-5f936abd7ae8 // indirect
	github.com/spf13/afero v1.15.0 // indirect
	github.com/spf13/cast v1.10.0 // indirect
	github.com/spf13/pflag v1.0.10 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	github.com/tmthrgd/go-hex v0.0.0-20190904060850-447a3041c3bc // indirect
	github.com/vmihailenco/msgpack/v5 v5.4.1 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	go.opentelemetry.io/otel v1.44.0 // indirect
	go.opentelemetry.io/otel/trace v1.44.0 // indirect
	go.yaml.in/yaml/v3 v3.0.5 // indirect
	go.yaml.in/yaml/v4 v4.0.0-rc.6 // indirect
	golang.org/x/exp v0.0.0-20260218203240-3dfff04db8fa // indirect
	golang.org/x/image v0.41.0 // indirect
	golang.org/x/mod v0.41.0 // indirect
	golang.org/x/sys v0.48.0 // indirect
	golang.org/x/text v0.42.0 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
	mellium.im/sasl v0.3.2 // indirect
	modernc.org/libc v1.68.0 // indirect
	modernc.org/mathutil v1.7.1 // indirect
	modernc.org/memory v1.11.0 // indirect
	modernc.org/sqlite v1.46.1 // indirect
	sigs.k8s.io/yaml v1.4.0 // indirect
)
