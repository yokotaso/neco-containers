import heapq
import mysql.connector as mc

HEAP_SIZE = 20

def _format_bytes(size):
    # 2**10 = 1024
    power = 2**10
    for unit in ('bytes', 'kb', 'mb', 'gb'):
       if size <= power:
           return "%d %s" % (size, unit)
       size /= power

    return "%d tb" % (size,)

def main():
    conn = mc.connect(unix_socket='/tmp/mysql.sock', user='root', password='cybozu')
    cur = conn.cursor()
    cur.execute('SHOW BINLOG EVENTS')
    size_heap = []
    for event in cur:
        info = event[5]
        if info.startswith('BEGIN'):
            start_log_pos = event[1]
        elif info.startswith('COMMIT '):
            end_log_pos = event[4]
            size = end_log_pos - start_log_pos
            if len(size_heap) < HEAP_SIZE:
                heapq.heappush(size_heap, size)
            else:
                heapq.heappushpop(size_heap, size)
    size_heap.sort(reverse=True)
    for size in size_heap:
        print(_format_bytes(size))
    cur.close()


if __name__ == '__main__':
    main()
