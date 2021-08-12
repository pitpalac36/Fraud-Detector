import React from 'react';
import './App.css';
import Transaction, {processData} from "./Transaction";
import headers from "./Headers";
import { Table } from 'antd';
import {MyPagination} from './MyPagination';
import * as dotenv from 'dotenv';



export class App extends React.Component<{}, { endpoint: string, transactions: Transaction[], columns:any }> {
    constructor(props: any) {
        super(props);
        this.state = {
            endpoint: "",
            transactions: [],
            columns: headers
        }
    }

    componentDidMount() {
        if(process.env.PRODUCTION !== "1") {
            dotenv.config();
        }
        this.setState({endpoint: process.env.REACT_APP_RESULTS_URL!});
        const ws = new WebSocket(process.env.REACT_APP_RESULTS_URL!);
        ws.onopen = () => {
        }
        ws.onmessage = e => {
            this.setState({
                transactions: this.state.transactions.concat(processData(JSON.parse(e.data)))
            })
        }
    }

    render(){
        return(
            <Table dataSource={this.state.transactions} columns={this.state.columns}>
            <MyPagination />
            </Table>
        )}
}
