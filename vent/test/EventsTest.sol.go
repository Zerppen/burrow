package test

import hex "github.com/tmthrgd/go-hex"

var Bytecode_EventsTest = hex.MustDecodeString("608060405234801561001057600080fd5b5061068d806100206000396000f300608060405260043610610062576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806306661abd146100675780638825519914610092578063c1de9c6d14610146578063dc667a6214610181575b600080fd5b34801561007357600080fd5b5061007c6101d4565b6040518082815260200191505060405180910390f35b34801561009e57600080fd5b506100cb6004803603810190808035906020019082018035906020019190919293919293905050506101dd565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561010b5780820151818401526020810190506100f0565b50505050905090810190601f1680156101385780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561015257600080fd5b5061017f6004803603810190808035906020019082018035906020019190919293919293905050506102a8565b005b34801561018d57600080fd5b506101d26004803603810190808035906020019082018035906020019190919293919293908035906020019082018035906020019190919293919293905050506103f8565b005b60008054905090565b6060600183836040518083838082843782019150509250505090815260200160405180910390206001018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561029b5780601f106102705761010080835404028352916020019161029b565b820191906000526020600020905b81548152906001019060200180831161027e57829003601f168201915b5050505050905092915050565b60006001838360405180838380828437820191505092505050908152602001604051809103902090508060020160009054906101000a900460ff16156103f35760008081548092919060019003919050555060018383604051808383808284378201915050925050509081526020016040518091039020600080820160006103309190610574565b6001820160006103409190610574565b6002820160006101000a81549060ff021916905550507f544553545f4556454e5453000000000000000000000000000000000000000000600019166103b684848080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050610566565b600019167fb64cbe0d18263bbda93ed76420a5e44f12291ff1187828c94336f06dcb61017860006040518082815260200191505060405180910390a35b505050565b60006001858560405180838380828437820191505092505050908152602001604051809103902090508060020160009054906101000a900460ff16151561044b5760008081548092919060010191905055505b848482600001919061045e9291906105bc565b5082828260010191906104729291906105bc565b5060018160020160006101000a81548160ff0219169083151502179055506104cb83838080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050610566565b600019167f544553545f4556454e54530000000000000000000000000000000000000000006000191661052f87878080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050610566565b600019167f6f50070fb9de82a81ea57052fbdf4459d17a1a9d68083b6f326b47bf17441e2960405160405180910390a45050505050565b600060208201519050919050565b50805460018160011615610100020316600290046000825580601f1061059a57506105b9565b601f0160209004906000526020600020908101906105b8919061063c565b5b50565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106105fd57803560ff191683800117855561062b565b8280016001018555821561062b579182015b8281111561062a57823582559160200191906001019061060f565b5b509050610638919061063c565b5090565b61065e91905b8082111561065a576000816000905550600101610642565b5090565b905600a165627a7a723058206be39370772ff802f524c05543e30774ffc73ec9d2c42c450777479461fc95b00029")
var Abi_EventsTest = []byte(`[{"constant":true,"inputs":[],"name":"count","outputs":[{"name":"size","type":"int256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"_name","type":"string"}],"name":"description","outputs":[{"name":"_description","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_name","type":"string"}],"name":"removeThing","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"_name","type":"string"},{"name":"_description","type":"string"}],"name":"addThing","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"anonymous":false,"inputs":[{"indexed":true,"name":"name","type":"bytes32"},{"indexed":true,"name":"key","type":"bytes32"},{"indexed":true,"name":"description","type":"bytes32"}],"name":"UpdateTestEvents","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"name","type":"bytes32"},{"indexed":true,"name":"key","type":"bytes32"},{"indexed":false,"name":"__DELETE__","type":"int256"}],"name":"DeleteTestEvents","type":"event"}]`)