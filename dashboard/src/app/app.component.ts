import {Component, OnDestroy, OnInit, ViewChild} from '@angular/core';
import {MatPaginator} from "@angular/material/paginator";

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit, OnDestroy {
  title = 'dashboard';
  displayedColumns: string[] = ['Source', 'Timestamp', 'V1', 'V2', 'V3', 'V4', 'V5', 'V6', 'V7', 'V8', 'V9', 'V10', 'V11', 'V12','V13', 'V14', 'V15', 'V16', 'V17', 'V18','V19', 'V20', 'V21', 'V22', 'V23', 'V24','V25', 'V26', 'V27', 'V28', 'Amount', 'Fraud']
  dataSource: any;

  @ViewChild(MatPaginator) paginator: MatPaginator | undefined;

  ngOnInit() {
    this.dataSource.paginator = this.paginator;
  }

  ngOnDestroy() {
  }
}
