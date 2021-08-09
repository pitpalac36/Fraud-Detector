import React from 'react';
import './App.css';
import Transaction, {processData} from "./Transaction";
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
            console.log(JSON.parse(e.data))
            this.setState({
                transactions: this.state.transactions.concat(processData(JSON.parse(e.data)))
            })
        }
    }

    itemRender(current: any, type: string, originalElement: any) {
        if (type === 'prev') {
            return <a>Previous</a>;
        }
        if (type === 'next') {
            return <a>Next</a>;
        }
        return originalElement;
    }

    render(){
        return(
            <Table dataSource={this.state.transactions} columns={this.state.columns} pagination={{itemRender: this.itemRender, pageSize: 50, style="padding: 24px"}}/>
        )}
}
