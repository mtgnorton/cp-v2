package shared

// 并发下创建和移除连接
//func Test_NewAndRemove(t *testing.T) {
//	slice := []string{"a", "b", "c", "d", "e"}
//	userIds := []uint{1, 2, 3, 4, 5}
//
//	ch := make(chan struct{}, 10)
//
//	m := garray.New(false)
//
//	wg := sync.WaitGroup{}
//
//	wg.Add(1000)
//
//	var count int64
//
//	gtimer.AddSingleton(gctx.New(), time.Second, func(ctx context.Context) {
//		fmt.Println(atomic.LoadInt64(&count))
//	})
//
//	for i := 0; i < 1000; i++ {
//
//		go func() {
//
//			ch <- struct{}{}
//
//			defer func() {
//				wg.Done()
//				<-ch
//			}()
//			userId := userIds[grand.N(0, 4)]
//
//			user := NewWsUser(nil, slice[grand.N(0, 4)], userId)
//			o := grand.N(0, 1)
//
//			if o == 0 {
//
//				index := WsService.AddUser(user)
//				m.Append([]int{int(userId), index})
//				fmt.Printf("+,id:%d,index:%d\n", userId, index)
//
//			} else {
//				itemInterface, found := m.PopRand()
//				if !found {
//					return
//				}
//				item := itemInterface.([]int)
//
//				fmt.Printf("-,id:%d,index:%d\n", item[0], item[1])
//
//				WsService.RemoveUser(uint(item[0]), 0)
//			}
//
//			atomic.AddInt64(&count, 1)
//
//		}()
//	}
//	wg.Wait()
//}
