import {Component, Inject} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';

export interface DialogData {
  clientSecret: string;
  title: string;
  displayedColumns: string[];
  data: string;
}

@Component({
  selector: 'app-dialog',
  templateUrl: './dialog-list.component.html',
})
export class DialogListComponent {
  public secret: string;
  public title: string;
  public displayedColumns: string[];
  public content: any;

  /**
   * Constructor.
   *
   * @param dialogRef MatDialogRef<DialogListComponent>
   * @param data DialogData
   */
  constructor(public dialogRef: MatDialogRef<DialogListComponent>,
              @Inject(MAT_DIALOG_DATA) public data: DialogData) {

    this.secret = data.clientSecret;
    this.title = data.title;
    this.displayedColumns = data.displayedColumns;
    this.content = data.data;
  }

  onClickSelection(key: string): void {
    this.dialogRef.close(key);
  }
}
