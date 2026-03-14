# birthday-bot
구글 캘린더에서 일정 조회 후 슬랙에 생일자 알림

프로젝트 초기화

```bash
go mod init <프로젝트 경로>
```

* `package main`은 진입점 패키지 (spring에서 `@SpringBootApplication` 있는 클래스)
* `func main()`은 `public static void main(String[] args)`

의존성 패키지 설치
`go get -u <package 경로>`

### 포인터 (*)

* `Config`는 값 자체를 복사
* `*Config`는 Config를 가리키는 포인터 (주소) -> 참조 타입처럼 동작
* `&Config`는 Config를 가리키는 주소를 가져옴

```go
// 값으로 전달 - 복사본이 넘어감
func updateValue(c Config) {
    c.Value = "new"  // 원본 안 바뀜
}

// 포인터로 전달 - 원본 주소가 넘어감
func updateValuePtr(c *Config) {
    c.Value = "new"  // 원본 바뀜
}

func main() {
    config := Config{Key: "test", Value: "old"}
    
    updateValue(config)
    fmt.Println(config.Value)  // "old" (안 바뀜)
    
    updateValuePtr(&config)    // &로 주소 전달
    fmt.Println(config.Value)  // "new" (바뀜!)
}
```

Go는 class가 없어서 생성자 주입을 아래와 같이함.

```kotlin
class ConfigRepository (
    private val db: Database
) {
    fun findByKey(Key: String): Config? {
        ...
    }
}

val repository = ConfigRepository(db)
```

```go
// struct로 필드만 정의 (class의 property)
type ConfigRepository struct {
    db *gorm.DB
}

// 팩토리 함수가 생성자 주입을 해주는 역할
func NewConfigRepository(db *gorm.DB) *ConfigRepository {
    return &ConfigRepository{db: db}
}

// 메서드는 struct에 붙이는 함수
func (r *ConfigRepository) FindByKey(key string) (*Config, error) {
    // r이 this 역할을 함
}
```