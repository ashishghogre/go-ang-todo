import { Component, OnInit } from '@angular/core';
import { Item } from './models/item';
import { ItemService } from './services/item.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit{
  items: Item[]

  constructor(private itemService: ItemService) {}

  ngOnInit() {
    this.getItems();
  }

  getItems(){
    this.itemService.getItems().subscribe((data: Item[]) => this.items = data);
  }

}
