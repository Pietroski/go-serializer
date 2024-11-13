// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	go_serializer "gitlab.com/pietroski-software-company/devex/golang/serializer/pkg/tools/serializer"
)

type FakeSerializer struct {
	DataRebindStub        func(interface{}, interface{}) error
	dataRebindMutex       sync.RWMutex
	dataRebindArgsForCall []struct {
		arg1 interface{}
		arg2 interface{}
	}
	dataRebindReturns struct {
		result1 error
	}
	dataRebindReturnsOnCall map[int]struct {
		result1 error
	}
	DeserializeStub        func([]byte, interface{}) error
	deserializeMutex       sync.RWMutex
	deserializeArgsForCall []struct {
		arg1 []byte
		arg2 interface{}
	}
	deserializeReturns struct {
		result1 error
	}
	deserializeReturnsOnCall map[int]struct {
		result1 error
	}
	SerializeStub        func(interface{}) ([]byte, error)
	serializeMutex       sync.RWMutex
	serializeArgsForCall []struct {
		arg1 interface{}
	}
	serializeReturns struct {
		result1 []byte
		result2 error
	}
	serializeReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSerializer) DataRebind(arg1 interface{}, arg2 interface{}) error {
	fake.dataRebindMutex.Lock()
	ret, specificReturn := fake.dataRebindReturnsOnCall[len(fake.dataRebindArgsForCall)]
	fake.dataRebindArgsForCall = append(fake.dataRebindArgsForCall, struct {
		arg1 interface{}
		arg2 interface{}
	}{arg1, arg2})
	stub := fake.DataRebindStub
	fakeReturns := fake.dataRebindReturns
	fake.recordInvocation("DataRebind", []interface{}{arg1, arg2})
	fake.dataRebindMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSerializer) DataRebindCallCount() int {
	fake.dataRebindMutex.RLock()
	defer fake.dataRebindMutex.RUnlock()
	return len(fake.dataRebindArgsForCall)
}

func (fake *FakeSerializer) DataRebindCalls(stub func(interface{}, interface{}) error) {
	fake.dataRebindMutex.Lock()
	defer fake.dataRebindMutex.Unlock()
	fake.DataRebindStub = stub
}

func (fake *FakeSerializer) DataRebindArgsForCall(i int) (interface{}, interface{}) {
	fake.dataRebindMutex.RLock()
	defer fake.dataRebindMutex.RUnlock()
	argsForCall := fake.dataRebindArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSerializer) DataRebindReturns(result1 error) {
	fake.dataRebindMutex.Lock()
	defer fake.dataRebindMutex.Unlock()
	fake.DataRebindStub = nil
	fake.dataRebindReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSerializer) DataRebindReturnsOnCall(i int, result1 error) {
	fake.dataRebindMutex.Lock()
	defer fake.dataRebindMutex.Unlock()
	fake.DataRebindStub = nil
	if fake.dataRebindReturnsOnCall == nil {
		fake.dataRebindReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.dataRebindReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSerializer) Deserialize(arg1 []byte, arg2 interface{}) error {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.deserializeMutex.Lock()
	ret, specificReturn := fake.deserializeReturnsOnCall[len(fake.deserializeArgsForCall)]
	fake.deserializeArgsForCall = append(fake.deserializeArgsForCall, struct {
		arg1 []byte
		arg2 interface{}
	}{arg1Copy, arg2})
	stub := fake.DeserializeStub
	fakeReturns := fake.deserializeReturns
	fake.recordInvocation("Deserialize", []interface{}{arg1Copy, arg2})
	fake.deserializeMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSerializer) DeserializeCallCount() int {
	fake.deserializeMutex.RLock()
	defer fake.deserializeMutex.RUnlock()
	return len(fake.deserializeArgsForCall)
}

func (fake *FakeSerializer) DeserializeCalls(stub func([]byte, interface{}) error) {
	fake.deserializeMutex.Lock()
	defer fake.deserializeMutex.Unlock()
	fake.DeserializeStub = stub
}

func (fake *FakeSerializer) DeserializeArgsForCall(i int) ([]byte, interface{}) {
	fake.deserializeMutex.RLock()
	defer fake.deserializeMutex.RUnlock()
	argsForCall := fake.deserializeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSerializer) DeserializeReturns(result1 error) {
	fake.deserializeMutex.Lock()
	defer fake.deserializeMutex.Unlock()
	fake.DeserializeStub = nil
	fake.deserializeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSerializer) DeserializeReturnsOnCall(i int, result1 error) {
	fake.deserializeMutex.Lock()
	defer fake.deserializeMutex.Unlock()
	fake.DeserializeStub = nil
	if fake.deserializeReturnsOnCall == nil {
		fake.deserializeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deserializeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSerializer) Serialize(arg1 interface{}) ([]byte, error) {
	fake.serializeMutex.Lock()
	ret, specificReturn := fake.serializeReturnsOnCall[len(fake.serializeArgsForCall)]
	fake.serializeArgsForCall = append(fake.serializeArgsForCall, struct {
		arg1 interface{}
	}{arg1})
	stub := fake.SerializeStub
	fakeReturns := fake.serializeReturns
	fake.recordInvocation("Serialize", []interface{}{arg1})
	fake.serializeMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSerializer) SerializeCallCount() int {
	fake.serializeMutex.RLock()
	defer fake.serializeMutex.RUnlock()
	return len(fake.serializeArgsForCall)
}

func (fake *FakeSerializer) SerializeCalls(stub func(interface{}) ([]byte, error)) {
	fake.serializeMutex.Lock()
	defer fake.serializeMutex.Unlock()
	fake.SerializeStub = stub
}

func (fake *FakeSerializer) SerializeArgsForCall(i int) interface{} {
	fake.serializeMutex.RLock()
	defer fake.serializeMutex.RUnlock()
	argsForCall := fake.serializeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSerializer) SerializeReturns(result1 []byte, result2 error) {
	fake.serializeMutex.Lock()
	defer fake.serializeMutex.Unlock()
	fake.SerializeStub = nil
	fake.serializeReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeSerializer) SerializeReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.serializeMutex.Lock()
	defer fake.serializeMutex.Unlock()
	fake.SerializeStub = nil
	if fake.serializeReturnsOnCall == nil {
		fake.serializeReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.serializeReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeSerializer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.dataRebindMutex.RLock()
	defer fake.dataRebindMutex.RUnlock()
	fake.deserializeMutex.RLock()
	defer fake.deserializeMutex.RUnlock()
	fake.serializeMutex.RLock()
	defer fake.serializeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSerializer) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ go_serializer.Serializer = new(FakeSerializer)
