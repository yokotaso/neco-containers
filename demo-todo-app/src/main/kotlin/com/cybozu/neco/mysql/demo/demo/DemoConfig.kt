package com.cybozu.neco.mysql.demo.demo

import org.springframework.boot.autoconfigure.jdbc.DataSourceProperties
import org.springframework.context.annotation.Bean
import org.springframework.context.annotation.Configuration

@Configuration
class DemoConfig {
    @Bean
    fun dataSource(properties: DataSourceProperties): MyDataSource? {
        return MyDataSource(properties.url, properties.username, properties.password)
    }
}