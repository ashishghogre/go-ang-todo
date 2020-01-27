import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import {NgbModule} from '@ng-bootstrap/ng-bootstrap'
import {HttpClientModule} from '@angular/common/http';

import { AppComponent } from './app.component';
import { HeaderComponent } from './components/header/header.component';
import { SideDrawerComponent } from './components/side-drawer/side-drawer.component';
import { TodoCardComponent } from './components/todo-card/todo-card.component';
import { BodyComponent } from './components/body/body.component';
import { EditableTodoItemComponent } from './components/editable-todo-item/editable-todo-item.component';
import { ReadOnlyTodoItemComponent } from './components/read-only-todo-item/read-only-todo-item.component';
import { EmptyTodoItemComponent } from './components/empty-todo-item/empty-todo-item.component';


@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    SideDrawerComponent,
    TodoCardComponent,
    BodyComponent,
    EditableTodoItemComponent,
    ReadOnlyTodoItemComponent,
    EmptyTodoItemComponent
  ],
  imports: [
    BrowserModule,
    NgbModule,
    HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
