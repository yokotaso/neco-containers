package com.cybozu.neco.mysql.demo.demo

import org.springframework.data.annotation.Id
import org.springframework.data.repository.CrudRepository
import org.springframework.web.bind.annotation.*
import java.time.LocalDateTime

@RestController
class TodoController(private val todoRepository: TodoRepository) {

    @GetMapping("/todo/{id}")
    fun get(@PathVariable id: Long): Todo? = todoRepository.findById(id).orElse(Todo())

    @PostMapping("/todo")
    fun create(@RequestBody todo: Todo) = todoRepository.save(todo)

    @DeleteMapping("/todo/{id}")
    fun delete(@PathVariable id: Long) = todoRepository.deleteById(id)
}

class Todo {
    @Id
    var todoId: Long? = null
    var todoTitle: String? = null
    var finished = false
    var createdAt: LocalDateTime? = null
}

interface TodoRepository : CrudRepository<Todo?, Long?> {
}