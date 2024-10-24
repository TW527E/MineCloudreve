[English Version](https://github.com/TW527E/MineCloudreve/blob/master/README.md)

<h1 align="center">
  <br>
  <a href="https://cloudreve.org/" alt="logo" ><img src="https://raw.githubusercontent.com/TW527E/frontend/master/public/static/img/logo192.png" width="150"/></a>
  <br>
  MineCloudreve
  <br>
</h1>

<h4 align="center">由 @TW527E 進行修改，支援了多家雲端儲存驅動的公有雲端硬碟系統。</h4>

<p align="center">
  <a href="https://github.com/TW527E/MineCloudreve/actions/workflows/test.yml">
    <img src="https://img.shields.io/github/actions/workflow/status/TW527E/MineCloudreve/test.yml?branch=master&style=flat-square"
         alt="GitHub Test Workflow">
  </a>
  <a href="https://codecov.io/gh/TW527E/MineCloudreve">
    <img src="https://img.shields.io/codecov/c/github/TW527E/MineCloudreve?style=flat-square">
  </a>
  <a href="https://goreportcard.com/report/github.com/TW527E/MineCloudreve">
      <img src="https://goreportcard.com/badge/github.com/TW527E/MineCloudreve?style=flat-square">
  </a>
  <a href="https://github.com/TW527E/MineCloudreve/releases">
    <img src="https://img.shields.io/github/v/release/TW527E/MineCloudreve?include_prereleases&style=flat-square" />
  </a>
  <a href="https://hub.docker.com/r/tw527e/minecloudreve">
     <img src="https://img.shields.io/docker/image-size/tw527e/minecloudreve?style=flat-square"/>
  </a>
</p>

<p align="center">
  <a href="https://cloudreve.org">主頁</a> •
  <a href="https://demo.cloudreve.org">演示站</a> •
  <a href="https://forum.cloudreve.org/">討論社區</a> •
  <a href="https://docs.cloudreve.org/">文件</a> •
  <a href="https://github.com/TW527E/MineCloudreve/releases">下載</a> •
  <a href="https://t.me/cloudreve_official">Telegram 群組</a> •
  <a href="#scroll-許可證">許可證</a>
</p>


![Screenshot](https://raw.githubusercontent.com/cloudreve/docs/master/images/homepage.png)

## :warning: 本項目注意事項

* 未經完整、專業測試，不建議用於實際生產環境
* 僅供交流學習使用，嚴禁用於非法目的，造成一切後果自負

## :sparkles: 特性

* :cloud: 支援本機、從機、七牛、阿里云 OSS、騰訊云 COS、又拍云、OneDrive (包括世紀互聯版) 、S3相容協議 作為儲存端
* :outbox_tray: 上傳/下載 支援客戶端直傳，支援下載限速
* 💾 可對接 Aria2 離線下載，可使用多個從機節點分擔下載任務
* 📚 線上 壓縮/解壓縮、多檔案打包下載
* 💻 覆蓋全部儲存策略的 WebDAV 協議支援
* :zap: 拖拽上傳、目錄上傳、流式上傳處理
* :card_file_box: 檔案拖拽管理
* :family_woman_girl_boy:   多使用者、使用者組、多儲存策略
* :link: 建立檔案、目錄的分享鏈接，可設定自動過期
* :eye_speech_bubble: 視訊、影象、音訊、 ePub 線上預覽，文字、Office 文件線上編輯
* :art: 自定義配色、黑暗模式、PWA 應用、全站單頁應用、國際化支援
* :rocket: All-In-One 打包，開箱即用
* 🌈 ... ...

## :hammer_and_wrench: 部署

下載適用於您目標機器操作系統、CPU架構的主程式，直接執行即可。

```shell
# 解壓程式包
tar -zxvf minecloudreve_VERSION_OS_ARCH.tar.gz

# 賦予執行許可權
chmod +x ./cloudreve

# 啟動 Cloudreve
./cloudreve
```
****
以上為最簡單的部署示例，您可以參考 [文件 - 起步](https://docs.cloudreve.org/) 進行更為完善的部署。

## :gear: 構建

自行構建前需要擁有 `Go 1.18↑`（越新越好）、`node.js v16.20`（或是使用下面的[替代方案]()）、`yarn`、`zip`, [goreleaser](https://goreleaser.com/intro/) 等必要依賴。

#### 安裝 goreleaser

```shell
go install github.com/goreleaser/goreleaser@latest
```

#### 克隆程式碼

```shell
git clone --recurse-submodules https://github.com/TW527E/MineCloudreve.git
```

#### 編譯專案

```shell
goreleaser build --clean --single-target --snapshot
```

## :alembic: 技術棧

* [Go](https://golang.org/) + [Gin](https://github.com/gin-gonic/gin)
* [React](https://github.com/facebook/react) + [Redux](https://github.com/reduxjs/redux) + [Material-UI](https://github.com/mui-org/material-ui)

## :scroll: 許可證

GPL V3

### 簡介
+ 🌩 支援多家雲端儲存的云盤系統
+ 基於 [3.8.3開源版本](https://github.com/TW527E/MineCloudreve/releases/tag/3.8.3) n次開發
+ 拉取主線最新版原始碼
+ 更新依賴至較新版本
+ 合併部分pr
   - [frontend#167](https://github.com/cloudreve/frontend/pull/167)
   - [backend#1911](https://github.com/TW527E/MineCloudreve/pull/1911)
   - [backend#1949](https://github.com/TW527E/MineCloudreve/pull/1949)
+ 修復部分已知Bug
+ 新增一些實用功能

### 編譯
+ 還是如果不需要修改前端，直接構建後端即可，前端包已預置
+ 前端
   - 環境：NodeJS v16.20 *
   - 進入 assets 目錄：`cd assets`
   - 安裝依賴：`yarn install` *
   - 構建靜態：`yarn build` *
   - 打包檔案：`bash pakstatics.sh`
   - (註：包管理器一定要用yarn，否則會報錯)
+ 後端
   - 環境：Golang >= 1.18，越新越好
   - 進入原始碼目錄
   - 構建程式：`go build -ldflags "-s -w" -tags "go_json" .`



### 使用

### 將 Cloudreve 官方社區版升級至 MineCloudreve 版
1.  將原有的社區版數據庫備份後（非必要，但為防止升級失敗，因此建議備份）
2.  在 *MineCloudreve 存在的目錄*或 *MineCloudreve 的 Docker 容器內* 下執行以下命令：
   ```
   ./cloudreve --database-script OSSToMineCloudreve
   ```