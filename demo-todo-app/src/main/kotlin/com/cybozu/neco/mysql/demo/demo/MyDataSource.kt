package com.cybozu.neco.mysql.demo.demo

import org.springframework.jdbc.datasource.AbstractDataSource
import org.springframework.jdbc.datasource.SmartDataSource
import java.sql.Connection
import java.sql.DriverManager

class MyDataSource(private val url: String, private val username: String, private val password: String)
    : AbstractDataSource(), SmartDataSource {

    var cache: Connection? = null;
    override fun shouldClose(con: Connection): Boolean = false

    override fun getConnection(): Connection? = getConnection(username, password)

    override fun getConnection(username: String?, password: String?): Connection? {
        if (this.cache == null) {
            this.cache = DriverManager.getConnection(url, username, password)
        }
        return this.cache
    }
}