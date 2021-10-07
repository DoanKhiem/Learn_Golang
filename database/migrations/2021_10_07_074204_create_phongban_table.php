<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

class CreatePhongbanTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('phongban', function (Blueprint $table) {
            $table->id();
            $table->string('name');
            $table->date('ngaythanhlap');
            $table->string('sdt');
            $table->string('diachi');
            $table->text('ghichu');
            $table->timestamps();
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('phongban');
    }
}
