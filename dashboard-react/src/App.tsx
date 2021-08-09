import React from 'react';
import './App.css';
import Transaction, {jsonToTransaction} from "./Transaction";
import headers from "./headers";
import { Table } from 'antd';

export class App extends React.Component<{}, { endpoint: string, transactions: Transaction[], columns:any }> {
    constructor(props: any) {
        super(props);
        this.state = {
            endpoint: "ws://localhost:8084/results",
            transactions: [],
            columns: headers
        }
    }

    componentDidMount() {
        const ws = new WebSocket(this.state.endpoint);
        ws.onopen = () => {
        }
        ws.onmessage = e => {
            console.log(e.data)
            this.setState({
                transactions: this.state.transactions.concat(jsonToTransaction(e.data))
            })
        }
    }

    render(){
        return(
            <Table dataSource={this.state.transactions} columns={this.state.columns}/>
        )}
}
