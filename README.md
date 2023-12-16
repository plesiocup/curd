## バックエンドのエンドポイント一覧

### 未実装

### やってる
`/movielinClick?movieid=num&userid=num`  
    evaluationを追加
    (ユーザーの好みをもとにおすすめするためのやつ)  
- ボタンが押されたタイミングでその映画の評価、countたちをプラス(いまあるupdateそのまま使えそう)  

ここからpython
- 評価の再計算、更新
- (このタイミングで協調フィルタリングのアルゴをトリガー)


### できた
`/createData`  
    画像URL、タイトル、説明、映画カテゴリ、再生時間、評価、評価した人数、公開年がhttp通信で送られて来るのでデータをデータベースに格納する。
- スクレイピングしたデータの保存に使う

`/getSearchedData?searchId=num`  
    searchIdのデータを返す。
- 検索結果の呼び出しに使う
    
`/getData?movieid=num`  
    idのデータを返す。（普通に）

### できた（仮）
goでの実装はおわり、pythonでの処理とどうつなぐか考える

`/userbasedRecommend?userid=num(ホーム画面)`  
    映画id、画像URL、タイトル、説明、映画カテゴリ、再生時間、評価、評価した人数、公開年をjson形式で返す。
    (ユーザーが興味を示しているカテゴリの映画をレコメンド)
- goでuserIDを受け取ってvectorの上位10こくらいを返す
- vectorの更新処理はpython,トリガーは評価ボタンが押されたとき

`/contentbasedRecommend?searchId=num&userid=num(映画詳細画面)`  
    映画id、画像URL、タイトル、説明、映画カテゴリ、再生時間、評価、評価した人数、公開年をjson形式で返す。（searchIdとそれに関連したデータ）
    (その映画と同じカテゴリの中で高評価なもの)
- pythonでsearchIdを受け取りDBから情報取得、その映画と同じカテゴリの映画情報を取得してから重み付けランダムの処理をして返す
### ききたいこと


### かいけつ
- userbase -> そのユーザーと似たユーザーの傾向を見ておすすめ
- contentbase -> そのカテゴリの中で重み付けランダム
