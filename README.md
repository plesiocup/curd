## バックエンドのエンドポイント一覧

### 未実装
`/userbasedRecommend?userid=num(ホーム画面)`
    映画id、画像URL、タイトル、説明、映画カテゴリ、再生時間、評価、評価した人数、公開年をjson形式で返す。

`/contentbasedRecommend?searchId=num&userid=num(映画詳細画面)`
    映画id、画像URL、タイトル、説明、映画カテゴリ、再生時間、評価、評価した人数、公開年をjson形式で返す。（searchIdとそれに関連したデータ）


### やってる
`/movielinClick?movieid=num&userid=num`
    evaluationを追加
    (ユーザーの好みをもとにおすすめするためのやつ)
    
### できた
`/createData`
    画像URL、タイトル、説明、映画カテゴリ、再生時間、評価、評価した人数、公開年がhttp通信で送られて来るのでデータをデータベースに格納する。

`/getSearchedData?searchId=num`
    searchIdのデータを返す。
    
`/getData?movieid=num`
    idのデータを返す。（普通に）

### ききたいこと
- userbase->他のユーザーの評価を元におすすめ（あなたと似たユーザー）
- contentbase->上に＋でカテゴリ情報を含める（この作品を見ているユーザー）
で合ってる？
- create時に含まれる評価データは既存のもの？アプリ内で集めたもの？
