# go-ddd-sample

This is sample application for like the DDD architecture.

## Design

* application
    * Write business logic
* domain
    * Define interface
        * repository interface for infrastructure
    * Define struct
        * Entity struct that represent mapping to data model
* infrastructure
    * Implements repository interface
    * Solves backend technical topics
        * e.x. message queue, persistence with RDB
* interfaces
    * Write HTTP handler and middleware

#### References:

* https://speakerdeck.com/mercari/ja-golang-package-composition-for-web-application-the-case-of-mercari-kauru
* http://pospome.hatenablog.com/entry/2017/10/11/023848
* https://medium.com/@timakin/go%E3%81%AE%E3%83%91%E3%83%83%E3%82%B1%E3%83%BC%E3%82%B8%E6%A7%8B%E6%88%90%E3%81%AE%E5%A4%B1%E6%95%97%E9%81%8D%E6%AD%B4%E3%81%A8%E7%8F%BE%E7%8A%B6%E7%A2%BA%E8%AA%8D-fc6a4369337
